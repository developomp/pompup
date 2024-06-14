package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "KiwiTalk",
		Desc: "KakaoTalk on Linux",
		Setup: func() {
			if wrapper.IsAppImageInstalled("kiwitalk") {
				return
			}

			if !wrapper.IsFileInDir(wrapper.InHome("Downloads"), "kiwitalk") {
				wrapper.DownloadFromGitHub("KiwiTalk", "KiwiTalk", "kiwi-talk_.*AppImage")
				return
			}
		},
		Reminders: []string{
			"Install KiwiTalk via gearlever",
		},
	})
}
