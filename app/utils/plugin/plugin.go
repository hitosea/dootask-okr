package plugin

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Configs struct {
	Code          string   `json:"code"`           //应用code
	Name          string   `json:"name"`           //应用名称
	Icon          string   `json:"icon"`           //应用名称
	Path          string   `json:"path"`           //应用路径
	PathIcon      string   `json:"path_icon"`      //应用icon路径
	DockerContent []byte   `json:"docker_content"` //docker_content
	NginxContent  []byte   `json:"nginx_content"`  //nginx_content
	Manifest      Manifest `json:"manifest"`
	Config        Config   `json:"config"`
}

type Manifest struct {
	Type        string   `yaml:"type" json:"type"`
	Tabs        []string `yaml:"tabs" json:"tabs"`
	Name        string   `yaml:"name" json:"name"`
	Description string   `yaml:"description" json:"description"`
	Author      string   `yaml:"author" json:"author"`
	AuthorUrl   string   `yaml:"authorUrl" json:"authorUrl"`
	Version     string   `yaml:"version" json:"version"`
}

type Config struct {
	Static          bool        `yaml:"static" json:"static"`
	Port            string      `yaml:"port" json:"port"`
	Host            string      `yaml:"host" json:"host,omitempty"`
	DockerContainer string      `yaml:"docker_container" json:"docker_container,omitempty"`
	MenusConfig     MenusConfig `yaml:"menus_config" json:"menus_config,omitempty"`
}

type MenusConfig struct {
	Menu             []MenuConfig `yaml:"menu" json:"menu"`
	AppMoreMenu      []MenuConfig `yaml:"app_more_menu" json:"app_more_menu"`
	HeadDropdownMenu []MenuConfig `yaml:"head_dropdown_menu" json:"head_dropdown_menu"`
	SettingPageMenu  []MenuConfig `yaml:"setting_page_menu" json:"setting_page_menu"`
}

type MenuConfig struct {
	Id        int    `yaml:"id" json:"id,omitempty"`
	Code      string `yaml:"code" json:"code"`
	Title     string `yaml:"title" json:"title"`
	Icon      string `yaml:"icon" json:"icon"`
	Path      string `yaml:"path" json:"path"`
	Authority string `yaml:"authority" json:"authority"`
	Platform  string `yaml:"platform" json:"platform"`
}

// 指定目录路径
var Dir = "plugins"

// 获取配置
func GetConfig() ([]*Configs, error) {
	var datas []*Configs

	// 获取目录下的所有子目录
	if dirs, err := getDirs(Dir); err != nil {
		return nil, err
	} else {
		// 处理每个子目录
		for _, d := range dirs {
			//
			var manifest Manifest
			if err := parsingYml(d, d+"/manifest.yml", &manifest); err != nil {
				return nil, err
			}
			//
			var config Config
			if err := parsingYml(d, d+"/config.yml", &config); err != nil {
				return nil, err
			}
			//
			dockerContent, _ := ioutil.ReadFile(d + "/docker/docker-compose.yml")
			nginxContent, _ := ioutil.ReadFile(d + "/nginx.conf")
			//
			datas = append(datas, &Configs{
				Code:          strings.Replace(d, Dir+"/", "", -1),
				Name:          manifest.Name,
				Path:          d,
				PathIcon:      d + "/icon",
				DockerContent: dockerContent,
				NginxContent:  nginxContent,
				Manifest:      manifest,
				Config:        config,
			})
		}
	}

	return datas, nil
}

// 获取指定目录下的所有一级子目录
func getDirs(dir string) ([]string, error) {
	var dirs []string
	// 打开目录
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	// 获取目录下的所有文件和子目录
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	// 遍历目录下的所有文件和子目录
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, filepath.Join(dir, file.Name()))
		}
	}
	return dirs, nil
}

// 解析yml
func parsingYml(dir string, file string, data interface{}) error {
	fileText, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	// 加载环境变量
	godotenv.Load(dir + "/.env")
	fileText = []byte(os.ExpandEnv(string(fileText)))
	// 解析 YAML 数据
	if err := yaml.Unmarshal(fileText, data); err != nil {
		return err
	}
	return nil
}
