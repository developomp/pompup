package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "GMetronome",
		Desc: "metronome",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.gitlab.dqpb.GMetronome")
		},
	})
}
