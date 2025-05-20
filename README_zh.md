# Gosearch

![Gosearch Logo](images/README/info.png)

[English](README.md) | [中文](README_zh.md)

Gosearch 是一个简单的命令行工具，允许你直接从终端搜索内容并在浏览器中打开结果。

## 功能特点

- 支持多个搜索引擎
- Google 作为默认搜索引擎
- 可通过环境变量配置默认搜索引擎
- 直接从终端执行搜索命令
- 快速轻量级
- 跨平台支持（Windows、macOS、Linux）

## 安装

### 方法 1：使用 go install（推荐）

```bash
go install github.com/Lily-404/search@latest
```

### 方法 2：手动安装

1. 确保已安装 Go 编程语言。
2. 克隆此仓库：
   ```bash
   git clone https://github.com/Lily-404/search.git
   cd search
   ```
3. 编译 Go 程序：
   ```bash
   go build -o search
   ```
4. 将生成的 `search` 可执行文件移动到你的 `PATH` 目录：

   对于 macOS/Linux：
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
search -s
```

![设置默认搜索引擎](images/README/set-default.png)

### 运行搜索命令

在终端中运行以下命令执行搜索：

```bash
search [engine_name] <search_content>
```

- `engine_name`：可选参数，指定搜索引擎。支持如下 engine name：

```
google | bing | baidu | github | chatgpt | perplexity | jike | bilibili | youtube | duckduckgo | brave | stackoverflow | wikipedia | reddit | x | medium | quora | linkedin | arxiv | scholar | zhihu | douban | v2ex | xiaohongshu | weibo
```

- `search_content`：必填参数，指定要搜索的内容。

### 示例

1. 使用默认搜索引擎：
   ```bash
   search hello world
   ```
   ![默认搜索](images/README/default-search.png)

2. 使用指定搜索引擎：
   ```bash
   search google hello world
   search zhihu golang
   search weibo 热搜
   search v2ex go mod
   search xiaohongshu 美食
   ```
   ![特定搜索引擎搜索](images/README/specific-search.png)

3. 搜索带有特殊字符的内容：
   ```bash
   search "how to use git"
   ```

## 使用场景

- 从终端快速进行网络搜索
- 查找开发者文档
- 视频内容搜索
- 代码仓库搜索
- AI 驱动的搜索
- 学术研究

## 贡献

欢迎贡献和报告问题！你可以：

1. Fork 此仓库
2. 创建你的特性分支（`git checkout -b feature/amazing-feature`）
3. 提交你的更改（`git commit -m 'Add some amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 开启一个 Pull Request

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## Star 历史

[![Star History Chart](https://api.star-history.com/svg?repos=Lily-404/search&type=Date)](https://star-history.com/#Lily-404/search&Date) 
