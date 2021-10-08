# 框架安装与项目构建步骤

## 1. 安装Wails

请参见链接：https://wails.app/zh/gettingstarted/linux/ 。

注意：`go get` 方法在 go version 1.17中已被废弃，
请不要使用教程中提供的 `go get -u github.com/wailsapp/wails/cmd/wails` 命令。


而是使用：`go install github.com/wailsapp/wails/cmd/wails@latest`

只需要执行到上述步骤。


## 2. 程序构建

在项目根目录运行 `wails build`（运行 `wails build -d`以 debug 模式构建），
构建后的可执行文件为 `./build/tool` 。