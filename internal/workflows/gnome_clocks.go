package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/dconf/gnome-clocks.conf
var _gnomeClocksDconf string

func init() {
	register(&Workflow{
		Name: "GNOME Clocks",
		Desc: "GNOME Time management utility",
		Tags: []Tag{Gnome, Gui},
		Setup: func() {
			wrapper.Flatpak("org.gnome.clocks")

			err := wrapper.Dconf(_gnomeClocksDconf)
			if err != nil {
				pterm.Fatal.Printfln("Failed to load config file: %s", err)
			}
		},
	})
}
