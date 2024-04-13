package installers

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Fonts",
		Desc: "fonts",
		Setup: func() {
			// https://github.com/google/fonts
			var fonts = [...]string{
				"ttf-ms-fonts",                    // MS fonts
				"adobe-source-han-sans-otc-fonts", // Korean font
				"ttf-baekmuk",                     // Korean font
				"unicode-emoji",                   // Colorful emoji
				"ttf-nerd-fonts-symbols-mono",     // Nerd font
				"ttf-d2coding-nerd",               // Korean coding font
				"noto-fonts",                      // cjk, emoji, etc
				"https://github.com/google/fonts/raw/main/ofl/audiowide/Audiowide-Regular.ttf",
				"https://github.com/google/fonts/raw/main/ofl/varelaround/VarelaRound-Regular.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosans/NotoSans%5Bwdth,wght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosans/NotoSans-Italic%5Bwdth,wght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosanskr/NotoSansKR%5Bwght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosansjp/NotoSansJP%5Bwght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosanstc/NotoSansTC%5Bwght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosanssc/NotoSansSC%5Bwght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosansmono/NotoSansMono%5Bwdth,wght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/nanumgothiccoding/NanumGothicCoding-Regular.ttf",
				"https://github.com/google/fonts/raw/main/ofl/nanumgothiccoding/NanumGothicCoding-Bold.ttf",
			}

			fontsDir := wrapper.InHome(".local/share/fonts")
			if !wrapper.PathExists(fontsDir) {
				err := os.MkdirAll(fontsDir, wrapper.DefaultDirPerm)
				if err != nil {
					pterm.Fatal.Printfln("Failed to create directory %s: %s", fontsDir, err)
				}
			}

			total := len(fonts)
			installed := 0
			for i, font := range fonts {
				if strings.HasPrefix(font, "https://") {
					fontFileName, _ := url.QueryUnescape(filepath.Base(font))
					fontPath := filepath.Join(fontsDir, fontFileName)

					if wrapper.PathExists(fontPath) {
						continue
					}

					pterm.Debug.Printfln("Installing font [%v / %v]: %s", i+1, total, font)
					err := wrapper.Run("wget", "-q", font, "-O", fontPath)
					if err != nil {
						pterm.Fatal.Printfln("Failed to install font %s to %s: %s", font, fontPath, err)
					}

					installed += 1
				} else {
					if wrapper.IsArchPkgInstalled(font) {
						continue
					}

					pterm.Debug.Printfln("Installing font [%v / %v]: %s", i+1, total, font)
					wrapper.Paru(font)
					installed += 1
				}
			}

			if installed == 0 {
				return
			}

			pterm.Debug.Println("Refreshing Font Cache")
			err := wrapper.Run("fc-cache", "-vf")
			if err != nil {
				pterm.Fatal.Println("Failed to regenerate font cache", err)
			}
		},
	})
}
