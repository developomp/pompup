package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Conda",
		Desc: "Python anaconda",
		Setup: func() {
			wrapper.ParuOnce("python-conda")
		},
	})
}
