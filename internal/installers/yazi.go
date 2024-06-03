package installers

import "github.com/developomp/pompup/internal/wrapper"

// https://github.com/sxyazi/yazi
func init() {
	register(&Installer{
		Name: "yazi",
		Desc: "CLI file manager",
		Setup: func() {
			wrapper.ParuOnce("yazi")
		},
	})
}
