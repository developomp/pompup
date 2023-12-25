package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "OpenTofu",
		Desc: "terraform but not evil",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("opentofu-bin")
		},
	})
}
