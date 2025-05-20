package search

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search content using the default search engine",
	Long:  `Search content using the default search engine and open the results in your browser.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

func init() {
	rootCmd.AddCommand(searchCmd)
}

func interactiveSearch() {
	prompt := promptui.Select{
		Label: "Select Search Engine",
		Items: engines,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U0001F449 {{ .Name | cyan }}",
			Inactive: "  {{ .Name | white }}",
			Selected: "\U0001F449 {{ .Name | green }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		color.Red("Prompt failed %v\n", err)
		return
	}

	selectedEngine := engines[index]

	queryPrompt := promptui.Prompt{
		Label: "Enter search query",
	}

	query, err := queryPrompt.Run()
	if err != nil {
		color.Red("Prompt failed %v\n", err)
		return
	}

	performSearch(selectedEngine.Name, query)
}

func performSearch(engineName string, query string) {
	var engine *Engine
	for _, e := range engines {
		if e.Name == engineName {
			engine = &e
			break
		}
	}

	if engine == nil {
		color.Red("Invalid search engine: %s", engineName)
		return
	}

	url := engine.URL
	if engine.Name == "jike" {
		url = fmt.Sprintf("https://web.okjike.com/search?keyword=%s&q=%s",
			strings.ReplaceAll(query, " ", "+"),
			strings.ReplaceAll(query, " ", "+"))
	} else if strings.Contains(url, "%s") {
		url = fmt.Sprintf(url, strings.ReplaceAll(query, " ", "+"))
	}

	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		color.Red("Unsupported platform")
		os.Exit(1)
	}

	if err != nil {
		color.Red("Error opening browser: %v", err)
		os.Exit(1)
	}
}
