package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Assembler",
		Desc: "Assembler",
		Setup: func() {
			wrapper.ParuOnce("yasm") // for x86_64, rewrite of nasm
		},
	})
}
