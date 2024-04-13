package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/dconf/gnome-clocks.conf
var _gnomeClocksDconf string

func init() {
	register(&Installer{
		Name: "GNOME Clocks",
		Desc: "GNOME Time management utility",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.clocks")
			wrapper.TryDconf(_gnomeClocksDconf)
		},
	})
}
