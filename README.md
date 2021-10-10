# 直接运行

项目目录中，`./build/tool_linux`为在Debian 10下构建的Linux可执行文件，可尝试直接运行。

在`./build`目录下，`chmod a+x tool_linux`赋予权限后即可运行。



对于macOS系统，双击运行`./build`目录下`Backup and Restore Tool.app`。

（在【系统偏好设置-隐私-完全磁盘访问】中添加该app，方可正常运行。）



# 框架安装与项目构建步骤

## 0. 注意

安装依赖时连接的服务器可能在海外，请耐心等候或给Terminal配置代理。

本文档针对Linux下的环境配置，若系统为macOS，请参见：https://wails.app/zh/gettingstarted/mac/。

**本项目不支持Windows系统！**

### 支持的Linux发行版

| Distro | Version                    |
| :----- | :------------------------- |
| Debian | 8, 9, 10                   |
| Ubuntu | 16.04, 18,04, 19.04, 19.10 |
| CentOS | 6, 7                       |
| Fedora | 29, 30                     |




## 1. 项目依赖安装

### a) Go

使用系统软件包管理器或从 [Go 下载页面](https://golang.org/dl/)下载并安装Go.

确保遵循官方的 [Go 安装说明](https://golang.org/doc/install#install).

将 `$GOPATH/bin`添加到 `PATH` 将 `on` 添加到 `GO111MODULE` 环境变量. 也可以将以下内容放到 `/etc/profile` (for a system-wide installation) or `$HOME/.profile`文件中:

```bash
export PATH=$PATH:$GOPATH/bin
export GO111MODULE=on
```

_注意：对配置文件的更改可能要等到下一次登录计算机后才能应用。 想要立即生效, 只需要运行如 `source $HOME/.profile`之类的 shell 命令即可_



### b) npm

从 [Node Downloads Page](https://nodejs.org/en/download/) 下载 `npm`.

运行 `npm --version` 验证安装是否成功.



### c) gcc, gtk, webkit

对于 Linux, Wails 使用 `gcc`, `webkit` and `GTK`.  这些需要使用下面的特定于发行版的命令进行安装.

#### Debian/Ubuntu 及其衍生版本

`sudo apt install build-essential libgtk-3-dev libwebkit2gtk-4.0-dev`

#### Arch Linux 及其衍生版本

`sudo pacman -S gcc pkgconf webkit2gtk gtk3`

#### Centos

`sudo yum install gcc-c++ make pkgconf-pkg-config webkitgtk3-devel gtk3-devel`

#### Fedora

`sudo yum install gcc-c++ make pkgconf-pkg-config webkit2gtk3-devel gtk3-devel`



### d) Wails

使用：`go install github.com/wailsapp/wails/cmd/wails@latest`

只需要执行到上述步骤。



可参见链接：https://wails.app/zh/gettingstarted/linux/ 。

注意：`go get` 方法在 go version 1.17中已被废弃，

请不要使用上述链接中提供的 `go get -u github.com/wailsapp/wails/cmd/wails` 命令。




## 2. 程序构建

终端切换到项目根目录，使用 `wails build`构建可执行文件（使用 `wails build -d`以 debug 模式构建）.

构建后的可执行文件为 `./build/tool` ，在终端中运行即可.