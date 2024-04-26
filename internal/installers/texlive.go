package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "texlive",
		Desc: "LaTeX document for true scholar",
		Setup: func() {
			wrapper.ParuOnce("texlive-meta")
		},
	})
}
