package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Librewolf",
		Desc: "privacy-focused firefox fork",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("io.gitlab.librewolf-community")
		},
	})
}
