package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Flatseal",
		Desc: "flatpak permission manager",
		Setup: func() {
			install.Flatpak("com.github.tchx84.Flatseal")
		},
	})
}
