package main

import (
	"dootask-okr/command"
	"runtime"
)

// @title DooTask
// @version 1.0
// @description  DooTask是一款轻量级的开源在线项目任务管理工具，提供各类文档协作工具、在线思维导图、在线流程图、项目管理、任务分发、即时IM，文件管理等工具。
// @termsOfService https://www.hitosea.com/
// @license.name AGPL-3.0 license
// @license.url http://www.gnu.org/licenses/
// @host http://localhost
// @BasePath /api/v1

//go:generate swag init -o ./app/api/v1/docs -g ./main.go -d ./app -g ../main.go
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	command.Execute()
}
