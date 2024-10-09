package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	engine := os.Getenv("DEFAULT_SEARCH_ENGINE")
	if engine == "" {
		engine = "google"
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
	case "chatgpt":
		url = fmt.Sprintf("https://chat.openai.com?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "perplexity":
		url = fmt.Sprintf("https://www.perplexity.ai/?q=%s", strings.ReplaceAll(query, " ", "+"))
	case "jike":
		url = fmt.Sprintf("https://web.okjike.com/search?keyword=%s", strings.ReplaceAll(query, " ", "+"))
	default:
		fmt.Println("This engine is not supported")
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
