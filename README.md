# Gosearch

![1744819750767](images/README/info.png)

[English](README.md) | [中文](README_zh.md)

Gosearch is a simple command-line tool that allows you to search content directly from your terminal and open the results in your browser. It supports multiple search engines, including Google, Bing, Baidu, Bilibili, Jike, Youtube, ChatGPT, and GitHub.

## Features

- Support for multiple search engines
- Google as the default search engine
- Configurable default search engine through environment variables
- Direct search commands from the terminal

## Installation

1. Make sure you have Go programming language installed.
2. Clone this repository or download the `search.go` file.
3. Compile the Go program:

   ```
   go build -o search search.go
   ```
4. Move the generated `search` executable to your `PATH` directory, for example `/usr/local/bin`:

   ```
   sudo mv search /usr/local/bin/
   ```

   Alternatively, you can place the `search` file in your `$HOME/bin` directory and ensure it's in your `PATH`:

   ```
   mkdir -p $HOME/bin
   mv search $HOME/bin/
   export PATH=$HOME/bin:$PATH
   ```

## Usage

### Setting the Default Search Engine

```
search -s || --set
```

### Running Search Commands

Run the following command in the terminal to perform a search:

```
search [search_engine] <search_content>
```

- `search_engine`: Optional parameter to specify the search engine. Supports `google`, `bing`, `baidu`, `jike`, `perplextity`, `bilibibli`, `youtube`, `github`, `chatgpt`.
- `search_content`: Required parameter specifying the content to search for.

### Examples

1. Search using the default search engine:

   ```
   search hello
   ```
2. Search using a specific search engine:

   ```
   search bing hello
   ```

## Supported Search Engines

- Google
- Bing
- Baidu
- GitHub
- Perplexity
- ChatGPT
- Jike
- Bilibili
- Youtube

## Contributing

Contributions and issue reports are welcome. Please submit Issues or Pull Requests on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
