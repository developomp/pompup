package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Lunar Client",
		Desc: "Competitive Minecraft Launcher",
		Setup: func() {
			wrapper.FlatpakOnce("com.lunarclient.LunarClient")
		},
	})
}
