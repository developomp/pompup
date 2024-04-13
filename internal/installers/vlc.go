package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "VLC",
		Desc: "video player",
		Setup: func() {
			wrapper.FlatpakOnce("org.videolan.VLC")
		},
	})
}
