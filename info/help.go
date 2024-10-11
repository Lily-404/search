package info

import (
	"flag"
	"fmt"
)

func Help() {

	helpFlag := flag.Bool("h", false, "显示帮助信息")
	helpLongFlag := flag.Bool("help", false, "显示帮助信息")
	versionFlag := flag.Bool("v", false, "版本打印")
	versionLongFlag := flag.Bool("version", false, "版本打印")
	engineFlag := flag.Bool("w", false, "查看当前搜索引擎")
	enginLongFlag := flag.Bool("what", false, "查看当前搜索引擎")

	engine := flag.String("engine", "google", "默认搜索引擎")

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
		fmt.Printf("当前的搜索引擎: %s\n", *engine)
		return
	}

}

func ShowHelp() {
	fmt.Println("GoSearch —— 搜索无界")
	fmt.Println(" ")
	fmt.Println("Usage:")
	fmt.Println("  gosearch <command> [arguments]")
	fmt.Println(" ")
	fmt.Println("The commands are:")
	fmt.Println("  -h, --help      显示帮助信息")
	fmt.Println("  -v, --version   版本打印")
	fmt.Println("  -w, --waht      查看当前的搜索引擎")
	fmt.Println("  -engine         设置默认搜索引擎")
	fmt.Println(" ")
	fmt.Println("Use \"gosearch help <command>\" for more information about a command.")
	fmt.Println(" ")
	fmt.Println("Additional help topics:")
	fmt.Println(" ")
	fmt.Println("  engine          搜索引擎选项")
	fmt.Println(" ")
	fmt.Println("Use \"gosearch help <topic>\" for more information about that topic.")
}
