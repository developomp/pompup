package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "Wireshark",
		Desc: "Packet analyzer",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("org.wireshark.Wireshark")
		},
	})
}
