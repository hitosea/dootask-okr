package command

import (
	"crypto/md5"
	"dootask-okr/config"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Message struct {
	ID    string
	Value string
}

// 使用有道翻译API翻译文本
type YoudaoResponse struct {
	ErrorCode   string   `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
}

func init() {
	rootCommand.AddCommand(translateCmd)
}

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "开始翻译",
	Run: func(_ *cobra.Command, _ []string) {
		godotenv.Load(".env")
		if config.YoudaoAppKey == "YOUDAO_APP_KEY" {
			config.YoudaoAppKey = os.Getenv("YOUDAO_APP_KEY")
		}
		if config.YoudaoAppSecret == "YOUDAO_SEC_KEY" {
			config.YoudaoAppSecret = os.Getenv("YOUDAO_SEC_KEY")
		}
		fmt.Printf("-------translate start-------\n")
		err := WriteLangFile()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("-------translate end-------\n")
	},
}

// WriteLangFile 将错误消息翻译并写入每种支持语言的 YAML 文件。
func WriteLangFile() error {
	// 定义常量
	const (
		i18nDir     = "i18n/lang" // i18n 目录
		fileMode    = 0644        // 文件权限
		defaultLang = "zh"        // 默认语言
	)

	// 获取支持的语言
	supportedLanguages := make([]language.Tag, len(config.Language))
	for i, lang := range config.Language {
		supportedLanguages[i] = language.Make(lang)
	}

	// 提取错误消息
	messages := extractMessages("app/constant/errs.go")

	// 并行翻译每种语言的消息
	var wg sync.WaitGroup
	for _, lang := range supportedLanguages {
		wg.Add(1)
		go func(lang language.Tag) {
			defer wg.Done()

			// 读取现有的翻译
			translatedData := make(map[string]string)
			filename := filepath.Join(i18nDir, lang.String()+".yaml")
			if _, err := os.Stat(filename); err == nil {
				translatedData = readYamlFile(filename)
			}

			// 翻译消息
			data := make(map[string]string)
			for _, message := range messages {
				// 检查消息是否已翻译
				if val, ok := translatedData[message.ID]; ok {
					data[message.ID] = val
					continue
				}
				// 不需要翻译的消息
				if lang.String() == defaultLang {
					data[message.ID] = message.Value
					continue
				}
				// 使用有道翻译 API 翻译消息
				language := lang.String()
				if language == "zh-Hant" {
					language = "zh-CHT"
				}
				translatedText, err := youdaoTranslateText(message.Value, "zh-CHS", language)
				if err != nil {
					log.Printf("Failed to translate message %q to %s: %v", message.ID, language, err)
					continue
				}
				data[message.ID] = translatedText
			}

			// 清空 YAML 文件
			file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, fileMode)
			if err != nil {
				log.Printf("Failed to open file %q: %v", filename, err)
				return
			}
			defer file.Close()
			if err := file.Truncate(0); err != nil {
				log.Printf("Failed to truncate file %q: %v", filename, err)
				return
			}

			// 将翻译写入文件
			if err := yaml.NewEncoder(file).Encode(data); err != nil {
				log.Printf("Failed to write translations to file %q: %v", filename, err)
			}
		}(lang)
	}

	// 等待所有翻译完成
	wg.Wait()

	return nil
}

// readYamlFile 读取 YAML 文件
func readYamlFile(filename string) map[string]string {
	data := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return data
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&data); err != nil {
		return data
	}

	return data
}

// 使用有道翻译API翻译文本
func youdaoTranslateText(text, from string, targetLang string) (string, error) {
	appKey := config.YoudaoAppKey
	appSecret := config.YoudaoAppSecret
	salt := strconv.FormatInt(time.Now().UnixNano(), 10)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(appKey+text+salt+appSecret)))

	if from == "" {
		from = "auto"
	}

	for {
		values := url.Values{}
		values.Set("q", text)
		values.Set("from", from)
		values.Set("to", targetLang)
		values.Set("appKey", appKey)
		values.Set("salt", salt)
		values.Set("sign", sign)

		resp, err := http.PostForm("http://openapi.youdao.com/api", values)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		var result YoudaoResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", err
		}

		if result.ErrorCode == "0" && len(result.Translation) > 0 {
			fmt.Println(result.Translation[0])
			return result.Translation[0], nil
		} else if result.ErrorCode == "411" {
			// 请求频率过快，等待一段时间后再次尝试
			time.Sleep(time.Second)
		} else {
			// 其他错误，直接返回
			return "", fmt.Errorf("translation failed with error code %s", result.ErrorCode)
		}
	}
}

// 从文件中提取消息
func extractMessages(filename string) []Message {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", source, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var messages []Message
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}

		for _, spec := range genDecl.Specs {
			valueSpec := spec.(*ast.ValueSpec)
			for _, name := range valueSpec.Names {
				comment := valueSpec.Comment
				if comment == nil {
					continue
				}

				msg := Message{
					ID: name.Name,
				}
				for _, c := range comment.List {
					if strings.HasPrefix(c.Text, "//") {
						msg.Value = strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
						break
					}
				}

				if msg.Value != "" {
					messages = append(messages, msg)
				}
			}
		}
	}

	return messages
}
