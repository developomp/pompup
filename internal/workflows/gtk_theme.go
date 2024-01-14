package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "pop! gtk theme",
		Desc: "GTK theme made by system76",
		Tags: []Tag{System, Gnome},
		Setup: func() {
			install.Paru("pop-gtk-theme")
		},
	})
}
