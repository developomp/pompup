package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
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
