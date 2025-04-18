# Gosearch

![1744819750767](images/README/info.png)

[English](README.md) | [中文](README_zh.md)

Gosearch是一个简单的命令行工具，可以在终端中直接搜索内容并打开浏览器进行搜索。支持多种搜索引擎，包括Google、Bing、baidu、bilibili、jike、Youtube、ChatGPT和GitHub等。

## 功能特性

- 支持多种搜索引擎
- 默认使用Google搜索引擎
- 可以通过环境变量设置默认搜索引擎
- 支持在终端中直接运行搜索命令

## 安装

1. 确保你已经安装了Go编程语言。
2. 克隆本仓库或下载`search.go`文件。
3. 编译Go程序：

   ```
   go build -o search search.go
   ```
4. 将生成的 `search` 可执行文件移动到你的 `PATH` 目录中，例如 `/usr/local/bin`：

   ```
   sudo mv search /usr/local/bin/
   ```

   或者，你可以将 `search` 文件放在 `$HOME/bin` 目录中，并确保该目录在你的 `PATH` 中：

   ```
   mkdir -p $HOME/bin
   mv search $HOME/bin/
   export PATH=$HOME/bin:$PATH
   ```

## 使用说明

### 设置默认搜索引擎

```
search -s || --set
```

### 运行搜索命令

在终端中运行以下命令进行搜索：

```
search [搜索引擎] <搜索内容>
```

- `搜索引擎`：可选参数，指定搜索引擎，支持 `google`、`bing`、`baidu`、`jike`、`perplextity`、`bilibibli`、`youtube`、`github`、`chatgpt`。
- `搜索内容`：必填参数，指定要搜索的内容。

### 示例

1. 使用默认搜索引擎进行搜索：

   ```
   search 你好
   ```
2. 指定搜索引擎进行搜索：

   ```
   search bing 你好
   ```

## 支持的搜索引擎

- Google
- Bing
- Baidu
- GitHub
- Perplexity
- ChatGPT
- Jike
- Bilibili
- Youtube

## 贡献

欢迎贡献代码和提出问题。请在 GitHub 上提交 Issue 或 Pull Request。

## 许可证

本项目采用 MIT 许可证。详细信息请参见 [LICENSE](LICENSE) 文件。 