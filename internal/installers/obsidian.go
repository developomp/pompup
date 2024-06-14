package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "Obsidian",
		Desc: "Best proprietary note taking app",
		Setup: func() {
			wrapper.FlatpakOnce("md.obsidian.Obsidian")
		},
	})
}
