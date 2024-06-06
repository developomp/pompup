package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Wireshark",
		Desc: "Packet analyzer",
		Setup: func() {
			wrapper.ParuOnce("wireshark-qt")
			wrapper.AddGroup("wireshark")
		},
	})
}
