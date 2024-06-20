package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "ImHex",
		Desc: "Hex Editor for Reverse Engineers",
		Setup: func() {
			wrapper.FlatpakOnce("net.werwolv.ImHex") // cspell:disable-line
		},
	})
}
