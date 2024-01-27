package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name:  "Golang",
		Desc:  "Golang Tools",
		Tags:  []Tag{Dev, Cli},
		Setup: setupGo,
	})
}

func setupGo() {
	install.Paru("go")
	install.Paru("go-tools")
}
