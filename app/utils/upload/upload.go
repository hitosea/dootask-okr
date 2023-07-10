package upload

import (
	"bytes"
	"crypto/md5"
	"dootask-okr/app/utils/common"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nfnt/resize"
)

// UploadParams 上传参数
type UploadParams struct {
	Type      string      // 上传类型
	File      io.Reader   // 上传文件
	Path      string      // 上传路径
	FileName  string      // 上传文件名
	Scale     []int       // 缩放尺寸 [压缩原图宽,高, 压缩方式]
	Size      int64       // 上传文件大小限制，单位KB
	AutoThumb bool        // 是否自动生成缩略图 false不要自动生成缩略图
	Chmod     os.FileMode // 文件权限 默认0644
	Compress  bool        // 是否压缩图片 默认true
}

// Image64Params image64图片保存参数
type Image64Params struct {
	Image64   string
	Path      string
	FileName  string
	Scale     []int
	AutoThumb bool
	Compress  bool
}

// UploadResult 上传结果
type UploadResult struct {
	Name   string  `json:"name"`   // 文件名
	Size   float64 `json:"size"`   // 文件大小，单位KB
	File   string  `json:"file"`   // 文件路径
	Path   string  `json:"path"`   // 文件路径
	URL    string  `json:"url"`    // 文件URL
	Thumb  string  `json:"thumb"`  // 缩略图URL
	Width  int     `json:"width"`  // 图片宽度
	Height int     `json:"height"` // 图片高度
	Ext    string  `json:"ext"`    // 文件扩展名
}

var (
	uploadRoot = "public" //文件根目录名称
)

// Upload 上传文件
func Upload(params UploadParams) (UploadResult, error) {
	params.Chmod = 0644
	params.Compress = true
	file := params.File
	if file == nil {
		return UploadResult{}, errors.New("您没有选择要上传的文件")
	}
	// 大小
	fileInfo, err := getFileSize(file)
	if err != nil {
		return UploadResult{}, err
	}
	fileSize := fileInfo.Size()
	if params.Size > 0 && fileSize > params.Size*1024 {
		return UploadResult{}, fmt.Errorf("文件大小超限，最大限制：%dKB", params.Size)
	}
	// 扩展名
	ext := getFileExt(params.FileName)
	lowerExt := strings.ToLower(ext)
	if !isAllowedType(params.Type, lowerExt) {
		return UploadResult{}, errors.New("文件格式错误，限制类型：" + strings.Join(getAllowedTypes(params.Type), ","))
	}
	// 路径
	err = os.MkdirAll(params.Path, os.ModePerm)
	if err != nil {
		return UploadResult{}, err
	}
	// 文件名
	md5sum, err := getFileMD5(file)
	if err != nil {
		return UploadResult{}, err
	}

	fileName := params.FileName
	if fileName == "" {
		fileName = fmt.Sprintf("%s.%s", md5sum, ext)
	}

	dstPath := filepath.Join(params.Path, fileName)
	os.Chmod(dstPath, params.Chmod)
	err = copySaveFile(file, dstPath)
	if err != nil {
		return UploadResult{}, err
	}

	uploadedFile, err := os.Open(dstPath)
	if err != nil {
		return UploadResult{}, err
	}
	defer uploadedFile.Close()

	var width, height int
	var thumbPath string
	// 图片时才进行压缩和缩略图处理
	if isImage(lowerExt) {
		// 压缩图片
		if params.Compress {
			if err := compressImage(dstPath, params.Scale); err != nil {
				fmt.Println(err)
				return UploadResult{}, errors.New("压缩图片失败")
			}
		}

		// 生成缩略图
		thumbPath = dstPath
		if params.AutoThumb {
			thumbPath = dstPath + "_thumb.jpg"
			if err := generateThumbnail(dstPath, thumbPath, 320, 0); err != nil {
				return UploadResult{}, errors.New("生成缩略图失败")
			}
		}

		// 图片获取宽高
		img, _, err := image.Decode(uploadedFile)
		if err != nil {
			return UploadResult{}, err
		}
		width = img.Bounds().Dx()
		height = img.Bounds().Dy()
	}

	dstPathAbs, _ := filepath.Abs(dstPath)          //绝对路径
	relativePath := TrimUploadsPath(dstPath)        //相对路径
	relativeThumbPath := TrimUploadsPath(thumbPath) //缩略相对路径
	return UploadResult{
		Name:   filepath.Base(fileName),
		Size:   float64(fileSize) / 1024.0,
		File:   dstPathAbs,
		Path:   relativePath,
		URL:    relativePath,      // 填写完整的URL
		Thumb:  relativeThumbPath, // 填写缩略图URL
		Width:  width,
		Height: height,
		Ext:    ext,
	}, nil
}

// Image64Save image64图片保存
func Image64Save(params Image64Params) (UploadResult, error) {
	params.Compress = true
	imgBase64 := params.Image64
	if !strings.HasPrefix(imgBase64, "data:image/") {
		return UploadResult{}, errors.New("图片格式错误")
	}

	// 解析图片格式和数据
	parts := strings.SplitN(imgBase64, ";base64,", 2)
	if len(parts) != 2 {
		return UploadResult{}, errors.New("图片格式错误")
	}
	ext := strings.TrimPrefix(parts[0], "data:image/")
	if !strings.Contains(ext, "/") {
		ext = "." + ext
	} else {
		ext = "." + strings.SplitN(ext, "/", 2)[1]
	}
	if !strings.Contains(ext, ".") {
		ext = "." + ext
	}
	if !strings.Contains(ext, "jpeg") && !strings.Contains(ext, "png") && !strings.Contains(ext, "gif") {
		return UploadResult{}, errors.New("图片格式错误")
	}
	imgData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return UploadResult{}, errors.New("图片格式错误")
	}

	// 生成文件名
	hash := md5.Sum(imgData)
	filename := fmt.Sprintf("paste_%x%s", hash, ext)
	if params.FileName != "" {
		filename = params.FileName
	}
	if len(params.Scale) == 3 {
		scaleName := fmt.Sprintf("%x_%dx%d_c%d%s", hash, params.Scale[0], params.Scale[1], params.Scale[2], ext)
		filename = scaleName
	}

	// 创建保存目录
	if err := os.MkdirAll(params.Path, os.ModePerm); err != nil {
		return UploadResult{}, errors.New("创建保存目录失败")
	}

	// 保存原图
	filepath := filepath.Join(params.Path, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return UploadResult{}, errors.New("保存原图失败")
	}
	defer file.Close()
	if _, err := io.Copy(file, bytes.NewReader(imgData)); err != nil {
		return UploadResult{}, errors.New("保存原图失败")
	}

	// 压缩图片
	if params.Compress {
		if err := compressImage(filepath, params.Scale); err != nil {
			return UploadResult{}, errors.New("压缩图片失败")
		}
	}

	// 生成缩略图
	thumbPath := filepath
	if params.AutoThumb {
		thumbPath = filepath + "_thumb.jpg"
		if err := generateThumbnail(filepath, thumbPath, 320, 0); err != nil {
			return UploadResult{}, errors.New("生成缩略图失败")
		}
	}

	// 获取图片信息
	imgFile, err := os.Open(filepath)
	if err != nil {
		return UploadResult{}, errors.New("获取图片信息失败")
	}
	defer imgFile.Close()
	img, _, err := image.DecodeConfig(imgFile)
	if err != nil {
		return UploadResult{}, errors.New("获取图片信息失败")
	}

	// 生成结果
	result := UploadResult{
		Name:   filename,
		Size:   float64(len(imgData)) / 1024,
		Path:   filepath,
		URL:    filepath,
		Thumb:  thumbPath,
		Width:  img.Width,
		Height: img.Height,
		Ext:    ext,
	}
	return result, nil
}

// 压缩图片
func compressImage(filepath string, scale []int) error {
	img, err := loadImage(filepath)
	if err != nil {
		return err
	}
	// 如果原始图片宽度小于压缩宽度，则按照原始宽高压缩
	if img.Bounds().Dx() < scale[0] {
		scale[0] = img.Bounds().Dx()
		scale[1] = img.Bounds().Dy()
	}
	// 小于10M的图片不压缩
	if img.Bounds().Dx()*img.Bounds().Dy() < 10000000 {
		return nil
	}
	resized := resizeImage(img, scale[0], scale[1])
	return saveImage(filepath, resized, 80)
}

// 生成缩略图
func generateThumbnail(filepath, thumbPath string, width, height int) error {
	img, err := loadImage(filepath)
	if err != nil {
		return err
	}

	resized := resizeImage(img, width, height)
	return saveImage(thumbPath, resized, 80)
}

// loadImage 加载图片文件并解码
func loadImage(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// saveImage 保存图片文件
func saveImage(filepath string, img image.Image, quality int) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
}

// ResizeImage 调整图片大小
func resizeImage(img image.Image, width, height int) image.Image {
	if width == 0 && height == 0 {
		return img
	}
	if width == 0 {
		width = img.Bounds().Dx() * height / img.Bounds().Dy()
	}
	if height == 0 {
		height = img.Bounds().Dy() * width / img.Bounds().Dx()
	}
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

// 获取文件大小
func getFileSize(file io.Reader) (os.FileInfo, error) {
	if fileInfo, ok := file.(os.FileInfo); ok {
		return fileInfo, nil
	}
	if seeker, ok := file.(io.Seeker); ok {
		pos, err := seeker.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, err
		}
		defer seeker.Seek(pos, io.SeekStart)
	}
	return os.Stat("/")
}

// 获取文件名扩展名
func getFileExt(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return ""
	}
	return strings.TrimPrefix(ext, ".")
}

// 获取文件MD5
func getFileMD5(file io.Reader) (string, error) {
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

// 复制保存文件
func copySaveFile(file io.Reader, dst string) error {
	// 将文件指针重置到文件开头
	if _, err := file.(io.ReadSeeker).Seek(0, 0); err != nil {
		return err
	}
	// 复制文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, file)
	return err
}

// 判断文件类型是否允许上传
func isAllowedType(uploadType, ext string) bool {
	types := getAllowedTypes(uploadType)
	return len(types) == 0 || contains(types, ext)
}

// 获取允许上传的文件类型
func getAllowedTypes(uploadType string) []string {
	switch uploadType {
	case "png":
		return []string{"png"}
	case "image":
		return []string{"jpg", "jpeg", "webp", "gif", "png"}
	case "video":
		return []string{"rm", "rmvb", "wmv", "avi", "mpg", "mpeg", "mp4"}
	case "audio":
		return []string{"mp3", "wma", "wav", "mid"}
	case "file":
		return []string{"txt", "doc", "xls", "ppt", "pdf", "zip", "rar"}
	default:
		return []string{}
	}
}

// 判断文件是否为图片
func isImage(ext string) bool {
	return contains([]string{"jpg", "jpeg", "webp", "gif", "png"}, ext)
}

// 判断字符串是否在字符串切片中
func contains(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// TrimUploadsPath 去除 uploads 前缀
func TrimUploadsPath(path string) string {
	// 将路径按照 / 分割成一个字符串切片
	pathParts := strings.Split(path, "/")
	// 找到 uploads 的位置
	uploadsIndex := -1
	for i, part := range pathParts {
		if part == "uploads" {
			uploadsIndex = i
			break
		}
	}
	// 如果找到了 uploads，将其后面的部分拼接成新的路径
	if uploadsIndex != -1 {
		newPath := strings.Join(pathParts[uploadsIndex:], "/")
		return newPath
	}
	return path
}

// Record64save base64语音保存
func Record64save(recordBase64 string, path string) (map[string]interface{}, error) {
	if !strings.HasPrefix(recordBase64, "data:audio/") {
		return nil, fmt.Errorf("语音格式错误")
	}
	res := regexp.MustCompile(`^data:audio/(\w+);base64,`).FindStringSubmatch(recordBase64)
	if len(res) != 3 {
		return nil, fmt.Errorf("语音格式错误")
	}
	extension := res[1]
	if extension != "mp3" && extension != "wav" {
		return nil, fmt.Errorf("语音格式错误")
	}
	data, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(recordBase64, res[0]))
	if err != nil {
		return nil, err
	}
	fileName := "record_" + fmt.Sprintf("%x", md5.Sum(data)) + "." + extension
	fileDir := path
	filePath := filepath.Join(uploadRoot, fileDir)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(filepath.Join(filePath, fileName), data, 0644); err != nil {
		return nil, err
	}
	fileSize := float64(len(data)) / 1024
	array := map[string]interface{}{
		"name": fileName,                                         //原文件名
		"size": fmt.Sprintf("%.2f", fileSize),                    //大小KB
		"file": filepath.Join(filePath, fileName),                //文件的完整路径
		"path": filepath.Join(fileDir, fileName),                 //相对路径
		"url":  common.FillUrl(filepath.Join(fileDir, fileName)), //完整的URL
		"ext":  extension,                                        //文件后缀名
	}
	return array, nil
}
