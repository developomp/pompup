package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME File Roller",
		Desc: "Compression and Decompression",
		Setup: func() {
			wrapper.FlatpakOnce("org.gnome.FileRoller")
		},
	})
}
