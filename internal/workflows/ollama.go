package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "ollama",
		Desc: "LLM made easy",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("ollama")
		},
	})
}
