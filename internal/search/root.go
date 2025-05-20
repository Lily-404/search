package search

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var version = "1.2.0"

func PrintVersion() {
	fmt.Println(`
░██████╗░░█████╗░  ░██████╗███████╗░█████╗░██████╗░░█████╗░██╗░░██╗
██╔════╝░██╔══██╗  ██╔════╝██╔════╝██╔══██╗██╔══██╗██╔══██╗██║░░██║
██║░░██╗░██║░░██║  ╚█████╗░█████╗░░███████║██████╔╝██║░░╚═╝███████║
██║░░╚██╗██║░░██║  ░╚═══██╗██╔══╝░░██╔══██║██╔══██╗██║░░██╗██╔══██║
╚██████╔╝╚█████╔╝  ██████╔╝███████╗██║░░██║██║░░██║╚█████╔╝██║░░██║
░╚═════╝░░╚════╝░  ╚═════╝░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝`)
	color.Cyan("go search version: %s", version)
}

var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "A command line tool for searching content from various search engines",
	Long: `A command line tool that allows you to search content directly from your terminal 
and open the results in your browser.

Supported search engines:
  google | bing | baidu | github | chatgpt | perplexity | jike | bilibili | youtube | duckduckgo | brave | stackoverflow | wikipedia | reddit | x | medium | quora | linkedin | arxiv | scholar | zhihu | douban | v2ex | xiaohongshu | weibo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			PrintVersion()
			return nil
		}

		setFlag, _ := cmd.Flags().GetBool("set")
		if setFlag {
			setDefaultEngine()
			return nil
		}

		// 如果没有参数，进入交互式搜索
		if len(args) == 0 {
			interactiveSearch()
			return nil
		}

		// 检查第一个参数是否是搜索引擎名称
		firstArg := args[0]
		if isEngineName(firstArg) && len(args) >= 2 {
			// 如果第一个参数是搜索引擎名称，使用指定的搜索引擎
			query := strings.Join(args[1:], " ")
			performSearch(firstArg, query)
			return nil
		}

		// 否则使用默认搜索引擎
		InitConfig()
		engine := DefaultEngine
		if engine == nil {
			InitConfig()
			engine = DefaultEngine
		}
		performSearch(*engine, strings.Join(args, " "))
		return nil
	},
	Args: cobra.ArbitraryArgs,
}

func isEngineName(name string) bool {
	for _, engine := range engines {
		if engine.Name == name {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.Flags().BoolP("set", "s", false, "Set default search engine")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")
}

func setDefaultEngine() {
	prompt := promptui.Select{
		Label: "Select Default Search Engine",
		Items: engines,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U0001F449 {{ .Description | cyan }}",
			Inactive: "  {{ .Description | white }}",
			Selected: "\U0001F449 {{ .Description | green }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		color.Red("Prompt failed %v\n", err)
		return
	}

	selectedEngine := engines[index]
	SetDefaultEngine(selectedEngine.Name)

	// 使用颜色显示设置结果
	color.Green("\n✓ Default search engine set to: %s", selectedEngine.Description)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
