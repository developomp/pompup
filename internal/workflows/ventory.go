package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Ventoy",
		Desc: "Bootable USB on steroid",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Paru("ventoy-bin")
		},
	})
}
