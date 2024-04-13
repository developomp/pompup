package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Librewolf",
		Desc: "privacy-focused firefox fork",
		Setup: func() {
			wrapper.FlatpakOnce("io.gitlab.librewolf-community")
		},
	})
}
