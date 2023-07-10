package common

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// Mkdir 创建目录
func Mkdir(path string, perm os.FileMode) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, perm)
		if err != nil {
			return
		}
		err = os.Chmod(path, perm)
		if err != nil {
			return
		}
	}
	return err
}

// ReadFile 读取文件
func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}

// WriteFile 保存文件（string）
func WriteFile(path string, content string) error {
	return WriteByte(path, []byte(content))
}

// WriteByte 保存文件（byte）
func WriteByte(path string, fileByte []byte) error {
	fileDir := filepath.Dir(path)
	if !Exists(fileDir) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(path, fileByte, 0666)
}

// AppendToFile 追加文件
func AppendToFile(path string, content string) error {
	fileDir := filepath.Dir(path)
	if !Exists(fileDir) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}

// 根据path获取文件名称
func PathToFileName(path string) string {
	return filepath.Base(path)
}

// 获取远程文件内容
func FileGetContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 写入文件内容
func FilePutContents(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}
