package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// 定义命令行参数
	engine, query, help := parseFlags()

	if help {
		printHelp()
		return
	}

	if query == "" {
		query = strings.Join(flag.Args(), " ")
	}

	if query == "" {
		fmt.Println("请提供搜索查询")
		return
	}

	if query == "version"{
		fmt.Println("
░██████╗░░█████╗░  ░██████╗███████╗░█████╗░██████╗░░█████╗░██╗░░██╗
██╔════╝░██╔══██╗  ██╔════╝██╔════╝██╔══██╗██╔══██╗██╔══██╗██║░░██║
██║░░██╗░██║░░██║  ╚█████╗░█████╗░░███████║██████╔╝██║░░╚═╝███████║
██║░░╚██╗██║░░██║  ░╚═══██╗██╔══╝░░██╔══██║██╔══██╗██║░░██╗██╔══██║
╚██████╔╝╚█████╔╝  ██████╔╝███████╗██║░░██║██║░░██║╚█████╔╝██║░░██║
░╚═════╝░░╚════╝░  ╚═════╝░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝")
		fmt.Println("0.1")
		return
	}

	if engine == "" {
		engine = "google"
	}

	url := buildSearchURL(engine, query)

	err := openBrowser(url)
	if err != nil {
		fmt.Println("打开浏览器时出错:", err)
	}
}

func parseFlags() (string, string, bool) {
	var engine string
	var query string
	var help bool

	flag.StringVar(&engine, "engine", os.Getenv("DEFAULT_SEARCH_ENGINE"), "搜索引擎 (google, bing, baidu, github, chatgpt, perplexity, jike)")
	flag.StringVar(&query, "query", "", "搜索查询")
	flag.BoolVar(&help, "help", false, "显示帮助信息")
	flag.Parse()

	return engine, query, help
}

func printHelp() {
	fmt.Println("Usage: search [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func buildSearchURL(engine, query string) string {
	switch engine {
	case "google":
		return fmt.Sprintf("https://www.google.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "bing":
		return fmt.Sprintf("https://www.bing.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "baidu":
		return fmt.Sprintf("https://www.baidu.com/s?wd=%s", strings.ReplaceAll(query, " ", "+"))
	case "github":
		return fmt.Sprintf("https://github.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "chatgpt":
		return fmt.Sprintf("https://chat.openai.com?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "perplexity":
		return fmt.Sprintf("https://www.perplexity.ai/?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "jike":
		return fmt.Sprintf("https://web.okjike.com/search?keyword=%s", strings.ReplaceAll(query, " ", "+"))
	default:
		fmt.Println("不支持的搜索引擎")
		os.Exit(1)
		return ""
	}
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("不支持的操作系统")
	}
	return err
}
