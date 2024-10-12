package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gosearch/config"
	"gosearch/info"
)

func isHelpFlag(flag string) bool {
	return strings.HasPrefix(flag, "-") || strings.HasPrefix(flag, "--")
}

func main() {

	// engine := os.Getenv("DEFAULT_SEARCH_ENGINE")
	config.Init()
	engine := config.DefaultEngine

	query := strings.Join(os.Args[1:], " ")
	var url string

	if len(os.Args) >= 3 {
		newEngine := os.Args[1]
		config.SetDefaultEngine(newEngine)
		query = strings.Join(os.Args[2:], " ")
	} else if len(os.Args) == 2 && isHelpFlag(os.Args[1]) {
		info.Help()
		return
	}

	if engine == nil {
		config.Init()
		engine = config.DefaultEngine
	}

	switch *engine {
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
		err = fmt.Errorf("Unsupported operating system")
	}

	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}