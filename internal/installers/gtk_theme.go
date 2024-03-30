package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "pop! gtk theme",
		Desc: "GTK theme made by system76",
		Tags: []Tag{System, Gnome},
		Setup: func() {
			wrapper.Paru("pop-gtk-theme")
		},
	})
}
