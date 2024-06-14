package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "nixfmt",
		Desc: "official nix formatter",
		Setup: func() {
			wrapper.NixOnce("github:NixOS/nixfmt")
		},
	})
}
