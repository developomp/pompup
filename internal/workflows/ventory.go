package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Ventoy",
		Desc: "Bootable USB on steroid",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.Paru("ventoy-bin")
		},
	})
}
