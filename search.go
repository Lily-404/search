package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: search [搜索引擎] <搜索内容>")
		fmt.Println("搜索引擎: google, bing, baidu, github, luxirity")
		fmt.Println("示例: search 你好")
		return
	}

	engine := os.Getenv("DEFAULT_SEARCH_ENGINE")
	if engine == "" {
		engine = "google" // 默认搜索引擎
	}

	query := strings.Join(os.Args[1:], " ")
	var url string

	if len(os.Args) >= 3 {
		engine = os.Args[1]
		query = strings.Join(os.Args[2:], " ")
	}

	switch engine {
	case "google":
		url = fmt.Sprintf("https://www.google.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "bing":
		url = fmt.Sprintf("https://www.bing.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "baidu":
		url = fmt.Sprintf("https://www.baidu.com/s?wd=%s", strings.ReplaceAll(query, " ", "+"))
	case "github":
		url = fmt.Sprintf("https://github.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "luxirity":
		url = fmt.Sprintf("https://search.luxirty.com/search?q=%s", strings.ReplaceAll(query, " ", "+"))
	default:
		fmt.Println("不支持的搜索引擎")
		return
	}

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

	if err != nil {
		fmt.Println("打开浏览器时出错:", err)
	}
}