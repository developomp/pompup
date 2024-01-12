package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "HotWire",
		Desc: "wireshark lite",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Flatpak("com.github.emmanueltouzery.hotwire")
		},
	})
}
