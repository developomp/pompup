package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Wireshark",
		Desc: "Packet analyzer",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Flatpak("org.wireshark.Wireshark")
		},
	})
}
