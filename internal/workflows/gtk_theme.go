package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "pop! gtk theme",
		Desc: "GTK theme made by system76",
		Tags: []Tag{System, Gnome},
		Setup: func() {
			wrapper.Paru("pop-gtk-theme")
		},
	})
}
