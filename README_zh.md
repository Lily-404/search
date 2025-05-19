# Gosearch

![1744819750767](images/README/info.png)

[English](README.md) | [中文](README_zh.md)

Gosearch 是一个简单的命令行工具，允许你直接从终端搜索内容并在浏览器中打开结果。它支持多个搜索引擎，包括 Google、Bing、百度、Bilibili、即刻、Youtube、ChatGPT 和 GitHub。

## 功能特点

- 支持多个搜索引擎
- Google 作为默认搜索引擎
- 可通过环境变量配置默认搜索引擎
- 直接从终端执行搜索命令
- 快速且轻量级
- 跨平台支持（Windows、macOS、Linux）

## 安装方法

### 方法一：使用 go install（推荐）

```bash
go install github.com/Lily-404/search@latest
```

### 方法二：手动安装

1. 确保已安装 Go 编程语言。
2. 克隆此仓库：
   ```bash
   git clone https://github.com/Lily-404/search.git
   cd search
   ```
3. 编译 Go 程序：
   ```bash
   go build -o search search.go
   ```
4. 将生成的 `search` 可执行文件移动到你的 `PATH` 目录：

   macOS/Linux 系统：
   ```bash
   sudo mv search /usr/local/bin/
   ```

   或者放置在你的 `$HOME/bin` 目录：
   ```bash
   mkdir -p $HOME/bin
   mv search $HOME/bin/
   export PATH=$HOME/bin:$PATH
   ```

## 使用方法

### 设置默认搜索引擎

```bash
search -s || --set
```

### 执行搜索命令

在终端中运行以下命令进行搜索：

```bash
search [搜索引擎] <搜索内容>
```

- `搜索引擎`：可选参数，指定搜索引擎。支持 `google`、`bing`、`baidu`、`jike`、`perplextity`、`bilibibli`、`youtube`、`github`、`chatgpt`、`x`。
- `搜索内容`：必需参数，指定要搜索的内容。

### 使用示例

1. 使用默认搜索引擎搜索：
   ```bash
   search hello world
   ```

2. 使用特定搜索引擎搜索：
   ```bash
   search bing hello world
   search github golang
   search youtube music
   ```

3. 搜索包含特殊字符的内容：
   ```bash
   search "如何使用 git"
   ```

## 支持的搜索引擎

- Google（默认）
- Bing
- 百度
- GitHub
- Perplexity
- ChatGPT
- 即刻
- Bilibili
- Youtube
- X (formerly Twitter)

## 使用场景

- 从终端快速进行网页搜索
- 查找开发文档
- 搜索视频内容
- 搜索代码仓库
- AI 驱动的搜索
- 学术研究

## 参与贡献

欢迎提交问题和贡献代码！你可以：

1. Fork 本仓库
2. 创建你的特性分支（`git checkout -b feature/amazing-feature`）
3. 提交你的更改（`git commit -m '添加一些很棒的功能'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证。详情请参见 [LICENSE](LICENSE) 文件。

## Star 历史

[![Star History Chart](https://api.star-history.com/svg?repos=Lily-404/search&type=Date)](https://star-history.com/#Lily-404/search&Date) 