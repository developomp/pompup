package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "KiwiTalk",
		Desc: "KakaoTalk on Linux",
		Setup: func() {
			wrapper.NixOnce("nixpkgs#kiwitalk")
		},
	})
}
