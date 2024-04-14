package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Conda",
		Desc: "Python anaconda",
		Setup: func() {
			wrapper.ParuOnce("python-conda")

			err := wrapper.Run("conda", "config", "--set", "auto_activate_base", "false")
			if err != nil {
				pterm.Fatal.Println("Failed to disable conda auto activation:", err)
			}
		},
	})
}
