package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "neovim",
		Desc: "nvim + NvChad, the best text editor (allegedly)",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("neovim")
			install.Paru("ripgrep")

			helper.Run("git", "clone", "https://github.com/NvChad/NvChad", helper.InHome(".config/nvim"), "--depth", "1")
		},
		Reminders: []string{
			"Install fonts",
		},
	})
}
