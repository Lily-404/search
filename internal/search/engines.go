package search

type Engine struct {
	Name        string
	Description string
	URL         string
}

var engines = []Engine{
	{
		Name:        "google",
		Description: "Google Search",
		URL:         "https://www.google.com/search?q=%s",
	},
	{
		Name:        "bing",
		Description: "Bing Search",
		URL:         "https://www.bing.com/search?q=%s",
	},
	{
		Name:        "baidu",
		Description: "Baidu Search",
		URL:         "https://www.baidu.com/s?wd=%s",
	},
	{
		Name:        "github",
		Description: "GitHub Search",
		URL:         "https://github.com/search?q=%s",
	},
	{
		Name:        "chatgpt",
		Description: "ChatGPT",
		URL:         "https://chat.openai.com?q=%s",
	},
	{
		Name:        "perplexity",
		Description: "Perplexity AI",
		URL:         "https://www.perplexity.ai/search?q=%s",
	},
	{
		Name:        "jike",
		Description: "Jike",
		URL:         "https://web.okjike.com/search?q=%s&type=POST",
	},
	{
		Name:        "bilibili",
		Description: "Bilibili",
		URL:         "https://search.bilibili.com/all?keyword=%s",
	},
	{
		Name:        "youtube",
		Description: "YouTube",
		URL:         "https://www.youtube.com/results?search_query=%s",
	},
	{
		Name:        "duckduckgo",
		Description: "DuckDuckGo",
		URL:         "https://duckduckgo.com/?q=%s",
	},
	{
		Name:        "brave",
		Description: "Brave Search",
		URL:         "https://search.brave.com/search?q=%s",
	},
	{
		Name:        "stackoverflow",
		Description: "Stack Overflow",
		URL:         "https://stackoverflow.com/search?q=%s",
	},
	{
		Name:        "wikipedia",
		Description: "Wikipedia",
		URL:         "https://en.wikipedia.org/wiki/Special:Search?search=%s",
	},
	{
		Name:        "reddit",
		Description: "Reddit",
		URL:         "https://www.reddit.com/search/?q=%s",
	},
	{
		Name:        "x",
		Description: "X (formerly Twitter)",
		URL:         "https://twitter.com/search?q=%s",
	},
	{
		Name:        "medium",
		Description: "Medium",
		URL:         "https://medium.com/search?q=%s",
	},
	{
		Name:        "quora",
		Description: "Quora",
		URL:         "https://www.quora.com/search?q=%s",
	},
	{
		Name:        "linkedin",
		Description: "LinkedIn",
		URL:         "https://www.linkedin.com/search/results/all/?keywords=%s",
	},
	{
		Name:        "arxiv",
		Description: "arXiv",
		URL:         "https://arxiv.org/search/?query=%s",
	},
	{
		Name:        "scholar",
		Description: "Google Scholar",
		URL:         "https://scholar.google.com/scholar?q=%s",
	},
	{
		Name:        "zhihu",
		Description: "知乎",
		URL:         "https://www.zhihu.com/search?type=content&q=%s",
	},
	{
		Name:        "douban",
		Description: "豆瓣",
		URL:         "https://www.douban.com/search?q=%s",
	},
	{
		Name:        "v2ex",
		Description: "V2EX",
		URL:         "https://www.sov2ex.com/?q=%s",
	},
	{
		Name:        "xiaohongshu",
		Description: "小红书",
		URL:         "https://www.xiaohongshu.com/search_result?keyword=%s",
	},
	{
		Name:        "weibo",
		Description: "微博",
		URL:         "https://s.weibo.com/weibo?q=%s",
	},
}
