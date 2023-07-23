package project_structs

type Project struct {
	Cominnek Cominnek `json:"cominnek"`
}

type Cominnek struct {
	Git Git `json:"git"`
}

type Git struct {
	GitConfig
	Flow    string    `json:"flow"` // flow: "git", "github" or "custom"
	Version string    `json:"version"`
	Extends GitConfig `json:"extends"`
}

type GitConfig struct {
	Branches []Branch `json:"branches"`
}

type Branch struct {
	Name   string       `json:"name"`
	Path   string       `json:"path"`
	Type   string       `json:"type"` // type: "production", "development" or "test"
	From   string       `json:"from"`
	To     []string     `json:"to"`
	Config BranchConfig `json:"config"`
}

type Action struct {
	Dispatcher string   `json:"dispatcher"`
	Args       []string `json:"args"`
}

type BranchConfig struct {
	Actions     []Action `json:"actions"`
	Description string   `json:"description"`
	Hidden      bool     `json:"hidden"`
	Unique      bool     `json:"unique"`
}
