package extras

type TreeItem struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type GitHubResponse struct {
	SHA       string     `json:"sha"`
	URL       string     `json:"url"`
	Tree      []TreeItem `json:"tree"`
	Truncated bool       `json:"truncated"`
}

type FlowCache struct {
	Flows      []string `json:"flows"`
	LastUpdate int64    `json:"last_update"`
}

type CacheFile struct {
	FlowCache FlowCache `json:"flow_cache"`
}
