package installers

import (
	"fmt"
	"net/url"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Fonts",
		Desc: "fonts",
		Tags: []Tag{System},
		Setup: func() {
			// install https://github.com/Crosse/font-install
			setupGo()
			wrapper.Run("go", "install", "github.com/Crosse/font-install@latest")

			wrapper.Flatpak("com.github.tchx84.Flatseal")

			wrapper.Paru("ttf-ms-fonts")                    // MS fonts
			wrapper.Paru("adobe-source-han-sans-otc-fonts") // Korean font
			wrapper.Paru("ttf-baekmuk")                     // Korean font
			wrapper.Paru("unicode-emoji")                   // Colorful emoji
			wrapper.Paru("ttf-nerd-fonts-symbols-mono")     // Nerd font
			wrapper.Paru("ttf-d2coding-nerd")               // Korean coding font
			wrapper.Paru("noto-fonts")                      // cjk, emoji, etc

			// fonts to download
			for _, fontName := range [...]string{
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
			} {
				fontURL := fmt.Sprintf("https://fonts.google.com/download?family=%s", url.QueryEscape(fontName))

				pterm.Debug.Println("Installing font:", fontURL)
				wrapper.Run("font-install", fontURL)
			}

			// regenerate font cache
			wrapper.Run("fc-cache -vf")
		},
	})
}
