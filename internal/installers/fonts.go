package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

var fonts = [...]string{
	"Audiowide",
	"Varela Round",
	"Noto Sans",
	"Noto Sans KR", // Korean
	"Noto Sans JP", // Japanese
	"Noto Sans TC", // Traditional Chinese
	"Noto Sans SC", // Simplified Chinese
	"Noto Sans Mono",
	"Nanum Gothic Coding",
	"Noto Emoji",
}

func init() {
	register(&Installer{
		Name: "Fonts",
		Desc: "fonts",
		Tags: []Tag{System},
		Setup: func() {
			wrapper.ParuOnce("ttf-ms-fonts")                    // MS fonts
			wrapper.ParuOnce("adobe-source-han-sans-otc-fonts") // Korean font
			wrapper.ParuOnce("ttf-baekmuk")                     // Korean font
			wrapper.ParuOnce("unicode-emoji")                   // Colorful emoji
			wrapper.ParuOnce("ttf-nerd-fonts-symbols-mono")     // Nerd font
			wrapper.ParuOnce("ttf-d2coding-nerd")               // Korean coding font
			wrapper.ParuOnce("noto-fonts")                      // cjk, emoji, etc

			// install https://github.com/lordgiotto/google-font-installer
			if !wrapper.IsBinInstalled("gfi") {
				setupNodejs()
				wrapper.Run("npm", "install", "-g", "google-font-installer")
			}

			total := len(fonts)
			for i, font := range fonts {
				pterm.Debug.Printfln("Installing font [%v / %v]: %s", i+1, total, font)
				err := wrapper.Run("gfi", "download", font, "--dest", wrapper.InHome(".local/share/fonts"))
				if err != nil {
					pterm.Fatal.Printfln("Failed to install font %s: %s", font, err)
				}
			}

			// regenerate font cache
			err := wrapper.Run("fc-cache", "-vf")
			if err != nil {
				pterm.Fatal.Println("Failed to regenerate font cache", err)
			}
		},
	})
}
