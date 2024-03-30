package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "HotWire",
		Desc: "wireshark lite",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("com.github.emmanueltouzery.hotwire")
		},
	})
}
