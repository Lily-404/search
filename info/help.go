package info

import (
	"flag"
	"fmt"

	"gosearch/config"
)

func Help() {

	helpFlag := flag.Bool("h", false, "显示帮助信息")
	helpLongFlag := flag.Bool("help", false, "显示帮助信息")
	versionFlag := flag.Bool("v", false, "版本打印")
	versionLongFlag := flag.Bool("version", false, "版本打印")
	engineFlag := flag.Bool("w", false, "查看当前搜索引擎")
	enginLongFlag := flag.Bool("what", false, "查看当前搜索引擎")
	setEngineFlag := flag.Bool("s", false, "设置默认搜索引擎")
	setEngineLongFlag := flag.Bool("set", false, "设置默认搜索引擎")
	// engine := flag.String("engine", "google", "默认搜索引擎")

	engine := config.DefaultEngine
	flag.Parse()

	if *helpFlag || *helpLongFlag {
		ShowHelp()
		return
	}

	if *versionFlag || *versionLongFlag {
		PrintVersion()
		return
	}

	if *engineFlag || *enginLongFlag {
		if engine != nil {
			fmt.Printf("当前的搜索引擎: %s\n", *engine)
		} else {
			fmt.Println("当前的搜索引擎未设置")
		}
		return
	}

	if *setEngineFlag || *setEngineLongFlag {
		var newEngine string
		fmt.Print("输入新的默认搜索引擎: ")
		fmt.Scanln(&newEngine)
		if newEngine != "" {
			config.SetDefaultEngine(newEngine)
			fmt.Printf("默认搜索引擎已设置为: %s\n", newEngine)
			return
		} else {
			fmt.Println("错误: 请提供一个有效的搜索引擎名称")
			return
		}
	}

}

func ShowHelp() {
	fmt.Println("GoSearch —— Think Different")
	fmt.Println(" ")
	fmt.Println("Usage:")
	fmt.Println("  gosearch <content>")
	fmt.Println("  gosearch [engine] <content>")
	fmt.Println(" ")
	fmt.Println("The commands are:")
	fmt.Println("  -h, --help      Display help information")
	fmt.Println("  -v, --version   Print version")
	fmt.Println("  -w, --what      View current search engine")
	fmt.Println("  -s, --set       Set default search engine")
	fmt.Println(" ")
	fmt.Println("Supported search engines:")
	fmt.Println("  google | bing | baidu | jike | perplexity | bilibili | youtube | github |chatgpt")

}
