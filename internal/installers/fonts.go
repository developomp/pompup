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
			var fonts = [...]string{
				// General
				"ttf-ms-fonts", // MS fonts
				"https://github.com/google/fonts/raw/main/ofl/audiowide/Audiowide-Regular.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosans/NotoSans%5Bwdth,wght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosans/NotoSans-Italic%5Bwdth,wght%5D.ttf",

				// Emoji
				"unicode-emoji", // Colored emojis
				"noto-fonts-emoji",

				// Chinese
				"https://github.com/google/fonts/raw/main/ofl/notosanstc/NotoSansTC%5Bwght%5D.ttf",
				"https://github.com/google/fonts/raw/main/ofl/notosanssc/NotoSansSC%5Bwght%5D.ttf",

				// Japanese
				"https://github.com/google/fonts/raw/main/ofl/notosansjp/NotoSansJP%5Bwght%5D.ttf",

				// Korean
				"adobe-source-han-sans-kr-fonts",
				"ttf-baekmuk",
				"https://github.com/google/fonts/raw/main/ofl/notosanskr/NotoSansKR%5Bwght%5D.ttf",

				// coding (monospaced)
				"https://github.com/google/fonts/raw/main/ofl/notosansmono/NotoSansMono%5Bwdth,wght%5D.ttf",
				"https://github.com/romkatv/powerlevel10k-media/raw/master/MesloLGS%20NF%20Regular.ttf", // https://github.com/romkatv/powerlevel10k?tab=readme-ov-file#fonts
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
