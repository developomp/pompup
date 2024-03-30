package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name:  "Golang",
		Desc:  "Golang Tools",
		Tags:  []Tag{Dev, Cli},
		Setup: setupGo,
	})
}

func setupGo() {
	wrapper.Paru("go")
	wrapper.Paru("go-tools")
}
