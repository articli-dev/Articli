# Articli

[![Release](https://github.com/k8scat/Articli/actions/workflows/release.yaml/badge.svg)](https://github.com/k8scat/Articli/actions/workflows/release.yaml)
[![GitHub Repo stars](https://img.shields.io/github/stars/k8scat/articli?style=social)](https://github.com/k8scat/Articli/stargazers)
[![GitHub watchers](https://img.shields.io/github/watchers/k8scat/articli?style=social)](https://github.com/k8scat/Articli/watchers)
[![star](https://gitee.com/k8scat/articli/badge/star.svg?theme=dark)](https://gitee.com/k8scat/articli/stargazers)
[![codecov](https://codecov.io/gh/k8scat/Articli/branch/main/graph/badge.svg?token=045FCRVF27)](https://codecov.io/gh/k8scat/Articli)

**Articli** is an Article CLI tool for managing content in multi platforms.

**Articli** 是一个可以管理多个平台内容的命令行工具，
通过解析 `Markdown` 文件内容以及调用平台接口，实现内容管理。

最终目标是基于 **本地文件** + **Git 代码仓** 管理所有的文章，
并且可以通过命令行操作以及 CI/CD，实现文章在各个平台的发布、更新等功能。
这样做的好处有：

- 数据安全，既发布到了第三方平台，又可以通过 **Git 代码仓**管理，避免因平台问题导致数据丢失
- 可以实现自动化，比如文章推送到自动在多个平台发布、更新
- 面向程序员的 CLI 工具，可以实现更多个性化的操作

为本项目点赞将鼓励作者继续完善下去，欢迎提出建议、Bug、PR。

## 支持的平台

### 平台文章管理

- [掘金](https://juejin.cn)
- [开源中国](https://oschina.net)
- [CSDN](https://csdn.net)

### Git 代码仓文件管理

- [GitHub](https://github.com)
- [极狐 GitLab](https://jihulab.com)

## 安装

### NPM

```shell
npm install -g @k8scat/articli
```

### Homebrew

```shell
# 添加 tap
brew tap k8scat/tap
# 安装
brew install acli

# 一条命令直接安装
brew install k8scat/tap/acli

# 后续升级
brew update
brew upgrade k8scat/tap/acli
```

### Docker

```shell
# 将配置文件的目录挂载到容器内
docker run \
  -it \
  --rm \
  -v $HOME/.config/articli:/root/.config/articli \
  k8scat/articli:latest \
  juejin auth login

# 升级
docker pull k8scat/articli:latest
```

### 二进制

Please download from the [releases page](https://github.com/k8scat/Articli/releases).

### 源码编译

```shell
git clone https://github.com/k8scat/articli.git
cd articli
make
```

## 文章模板

我们将使用文件内容开头 `---` 之间的数据作为文章的配置信息（元数据），
根据配置信息在不同平台上创建或更新文章，参考 [文章模板](https://raw.githubusercontent.com/articli-dev/Articli/templates/article.md)。

## 使用说明

所有的命令都可以通过 `-h` 或 `--help` 参数查看帮助信息。

```shell
$ acli --help
Manage content in multi platforms.

Usage:
  acli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  csdn        Manage content in csdn.net
  github      Manage content in github.com
  help        Help about any command
  juejin      Manage content in juejin.cn
  oschina     Manage content in oschina.net
  version     Show version information

Flags:
  -c, --config string   An alternative config file
  -h, --help            help for acli

Use "acli [command] --help" for more information about a command.
```

### 查看版本

```shell
acli version
```

### 掘金

#### 登录

使用浏览器 Cookie 进行登录

```shell
# 交互式登录
acli juejin auth login

# 从标准输入获取 Cookie
acli juejin auth login --with-cookie < cookie.txt
```

#### 创建/更新文章

```shell
# create 命令可以通过识别文章的配置信息，自动选择创建或者更新文章，同时发布到掘金
acli juejin article create /path/to/article.md
```

#### 查看文章列表

通过 `-k` 或 `--keyword` 关键字参数过滤文章列表

```shell
acli juejin article list -k Docker
```

#### 打开文章

使用默认浏览器打开文章

```shell
acli juejin article view 7055689358657093646
```

#### 查看分类

```shell
acli juejin category list
```

#### 查看标签

```shell
# 过滤关键字
acli juejin tag list -k Go
```

#### 缓存标签

由于标签的数量比较多，可以通过设置缓存加快读取速度

```shell
# 设置缓存
acli juejin tag cache

# 使用缓存
acli jujin tag list --use-cache
```

#### 上传图片

支持上传本地图片和网络图片

```shell
# 本地图片
acli juejin image upload leetcode-go.png

# 网络图片
acli juejin image upload https://launchtoast.com/wp-content/uploads/2021/11/learn-rust-programming-language.png
```

### 开源中国

#### 登录

```shell
# 交互式登录
acli oschina auth login

# 从标准输入中读取 cookie
acli oschina auth login --with-cookie < cookie.txt
```

#### 创建/更新文章

```shell
acli oschina article create /path/to/article.md
```

### CSDN

#### 登录

```shell
# 交互式登录
acli csdn auth login

# 从标准输入中读取 cookie
acli csdn auth login --with-cookie < cookie.txt
```

#### 创建/更新文章

```shell
acli csdn article create /path/to/article.md
```

### GitHub

#### 登录

使用 [Personal Access Token](https://github.com/settings/tokens) 进行登录

```shell
# 交互式登录
acli github auth login

# 从标准输入获取 Token
acli github auth login --with-token < token.txt
```

#### 上传文件

```shell
# 上传本地文件
acli github file upload -o <owner> -r <repo> [-p <store path>] <local path>

# 上传网络资源
acli github file upload -o <owner> -r <repo> [-p <store path>] <resource url>
```

#### 列取文件

```shell
# 获取代码仓根目录的文件列表，包括文件和目录
acli github file get -o <owner> -r <repo>

# 指定 path
acli github file get -o <owner> -r <repo> -p <path>
```

#### 删除文件

```shell
# 支持同时删除多个文件
acli github file delete -o <owner> -r <repo> <path ...>
```

### 极狐 GitLab

#### 登录

使用 [Personal Access Token](https://jihulab.com/-/profile/personal_access_tokens) 进行登录

```shell
# 交互式登录，默认登录 https://jihulab.com
acli gitlab auth login

# 从标准输入获取 Token
acli gitlab auth login --with-token < token.txt

# 指定 GitLab 实例地址
acli gitlab auth login --base-url https://gitlab.com
```

#### 上传文件

```shell
# 上传本地文件
acli gitlab file upload -o <owner> -r <repo> [-p <store path>] <local path>

# 上传网络资源
acli gitlab file upload -o <owner> -r <repo> [-p <store path>] <resource url>

# 使用 projectID 代替 owner/repo
acli gitlab file upload --project-id <projectID> [-p <store path>] <local path>
```

#### 列取文件

```shell
# 获取代码仓根目录的文件列表，包括文件和目录
acli gitlab file get -o <owner> -r <repo>

# 指定 path
acli gitlab file get -o <owner> -r <repo> -p <path>
```

#### 删除文件

```shell
# 支持同时删除多个文件
acli github file delete -o <owner> -r <repo> <path ...>
```

### 简化命令

使用 `alias` 别名进行简化命令

```shell
# 将 acli juejin 简化成 jcli
cat >> ~/.bashrc << EOF
alias jcli="acli juejin"
alias ocli="acli oschina"
alias gcli="acli github"
EOF

# 生效
source ~/.bashrc

# 使用简化后的命令查看掘金的登录状态
jcli auth status
```
