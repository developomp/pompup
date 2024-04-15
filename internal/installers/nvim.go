package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "neovim",
		Desc: "nvim + NvChad, the best text editor (allegedly)",
		Setup: func() {
			wrapper.ParuOnce("neovim")
			wrapper.ParuOnce("ripgrep")

			configPath := wrapper.InHome(".config/nvim")
			if !wrapper.PathExists(configPath) {
				wrapper.Run("git", "clone", "https://github.com/NvChad/NvChad", configPath, "--depth", "1")
			}
		},
	})
}
