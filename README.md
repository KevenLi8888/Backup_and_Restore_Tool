# UI框架安装和项目构建步骤

## 0. 重要

前后端整合的过程中修改了小部分后端代码，
建议先在原来单独的后端调试好后再把修改部份合并入此项目，
或者把修改后的版本提供给 @Keven Li 去处理合并问题。

后端文件在./lib文件中，
在函数定义上稍有调整（去掉了main函数，变成package中的函数用于前端调用）。

目前主要逻辑和错误处理已经实现，未实现自定义备份路径和文件名功能。

## 1. 安装Wails

请参见链接：https://wails.app/zh/gettingstarted/linux/ 。

注意：`go get` 方法在 go version 1.17中已被废弃，
请不要使用教程中提供的 `go get -u github.com/wailsapp/wails/cmd/wails` 命令。


而是使用：`go install github.com/wailsapp/wails/cmd/wails@latest`

只需要执行到上述步骤。


## 2. 程序构建

在项目根目录运行 `wails build -d`（以 debug 模式构建），
构建后的可执行文件为 `./build/gui1` 。