package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "neovim",
		Desc: "nvim + NvChad, the best text editor (allegedly)",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("neovim")
			wrapper.Paru("ripgrep")

			wrapper.Run("git", "clone", "https://github.com/NvChad/NvChad", wrapper.InHome(".config/nvim"), "--depth", "1")
		},
		Reminders: []string{
			"Install fonts",
		},
	})
}
