package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Rust",
		Desc: "Rustup and stuff",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			if wrapper.IsArchPkgInstalled("rust") {
				// rust conflicts with rustup. Remove before installing rustup.

				if err := wrapper.Run("paru", "-R", "--noconfirm", "rust"); err != nil {
					pterm.Fatal.Println("Failed to remove rust package:", err)
				}
			}

			wrapper.ParuOnce("rustup")

			wrapper.Run("rustup", "install", "stable")
			wrapper.Run("rustup", "default", "stable")
		},
	})
}
