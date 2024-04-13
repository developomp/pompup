package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "DaVinci Resolve",
		Desc: "Video Editing Tool",
		Setup: func() {
			wrapper.ParuOnce("opencl-nvidia")
			wrapper.ParuOnce("noto-fonts")
			wrapper.ParuOnce("jdk-openjdk")
			wrapper.ParuOnce("davinci-resolve")
		},
	})
}
