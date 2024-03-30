package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "HotWire",
		Desc: "wireshark lite",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("com.github.emmanueltouzery.hotwire")
		},
	})
}
