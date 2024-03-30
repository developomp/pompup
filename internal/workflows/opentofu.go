package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "OpenTofu",
		Desc: "terraform but not evil",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("opentofu-bin")
		},
	})
}
