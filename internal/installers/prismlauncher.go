package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Prism Launcher",
		Desc: "Better Minecraft Launcher",
		Setup: func() {
			wrapper.FlatpakOnce("org.prismlauncher.PrismLauncher")
		},
	})
}
