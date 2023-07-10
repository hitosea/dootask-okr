package template

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Template 从模板中获取内容
func Template(original string, envMap map[string]any) string {
	tmpl, err := template.New("text").Parse(original)
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("模板分析失败: %s", err))
		}
	}()
	if err != nil {
		panic(1)
	}
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	content := string(buffer.Bytes())
	content = strings.ReplaceAll(content, "<no value>", "")
	return content
}
