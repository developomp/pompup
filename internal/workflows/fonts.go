package workflows

import (
	"fmt"
	"net/url"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Fonts",
		Desc: "fonts",
		Tags: []Tag{System},
		Setup: func() {
			// install https://github.com/Crosse/font-install
			setupGo()
			helper.Run("go", "install", "github.com/Crosse/font-install@latest")

			install.Flatpak("com.github.tchx84.Flatseal")

			install.Paru("ttf-ms-fonts")                    // MS fonts
			install.Paru("adobe-source-han-sans-otc-fonts") // Korean font
			install.Paru("ttf-baekmuk")                     // Korean font
			install.Paru("unicode-emoji")                   // Colorful emoji
			install.Paru("ttf-nerd-fonts-symbols-mono")     // Nerd font
			install.Paru("ttf-d2coding-nerd")               // Korean coding font
			install.Paru("noto-fonts")                      // cjk, emoji, etc

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
				helper.Run("font-install", fontURL)
			}

			// regenerate font cache
			helper.Run("fc-cache -vf")
		},
	})
}
