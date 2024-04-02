package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Wireshark",
		Desc: "Packet analyzer",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.FlatpakOnce("org.wireshark.Wireshark")
		},
	})
}
