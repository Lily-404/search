# Gosearch

![Gosearch Logo](images/README/info.png)

[English](README.md) | [中文](README_zh.md)

Gosearch is a simple command-line tool that allows you to search content directly from your terminal and open the results in your browser. 

## Features

- Support for multiple search engines
- Google as the default search engine
- Configurable default search engine through environment variables
- Direct search commands from the terminal
- Fast and lightweight
- Cross-platform support (Windows, macOS, Linux)

## Installation

### Method 1: Using go install (Recommended)

```bash
go install github.com/Lily-404/search@latest
```

### Method 2: Manual Installation

1. Make sure you have Go programming language installed.
2. Clone this repository:
   ```bash
   git clone https://github.com/Lily-404/search.git
   cd search
   ```
3. Compile the Go program:
   ```bash
   go build -o search
   ```
4. Move the generated `search` executable to your `PATH` directory:

   For macOS/Linux:
   ```bash
   sudo mv search /usr/local/bin/
   ```

   Or place it in your `$HOME/bin` directory:
   ```bash
   mkdir -p $HOME/bin
   mv search $HOME/bin/
   export PATH=$HOME/bin:$PATH
   ```

## Usage

### Setting the Default Search Engine

```bash
search -s
```

![Setting Default Engine](images/README/set-default.png)

### Running Search Commands

Run the following command in the terminal to perform a search:

```bash
search [engine_name] <search_content>
```

- `engine_name`:  Optional parameter to specify the search engine. The following engine names are supported:

```
google | bing | baidu | github | chatgpt | perplexity | jike | bilibili | youtube | duckduckgo | brave | stackoverflow | wikipedia | reddit | x | medium | quora | linkedin | arxiv | scholar | zhihu | douban | v2ex | xiaohongshu | weibo
```

- `search_content`: Required parameter specifying the content to search for.

### Examples

1. Use the default search engine:
   ```bash
   search hello world
   ```
   ![Default Search](images/README/default-search.png)

2. Use a specified search engine:
   ```bash
   search google hello world
   search zhihu golang
   search weibo 热搜
   search v2ex go mod
   search xiaohongshu 美食
   ```
   ![Specific Engine Search](images/README/specific-search.png)

3. Search for content containing special characters:
   ```bash
   search "how to use git"
   ```

## Use Cases

- Quick web searches from terminal
- Developer documentation lookup
- Video content search
- Code repository search
- AI-powered search
- Academic research

## Contributing

Contributions and issue reports are welcome! Please feel free to:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=Lily-404/search&type=Date)](https://star-history.com/#Lily-404/search&Date)
