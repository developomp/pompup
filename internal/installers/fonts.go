package installers

import (
	"net/url"
	"path/filepath"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

// https://github.com/google/fonts
var fontURLs = [...]string{
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

			total := len(fontURLs)
			for i, fontURL := range fontURLs {
				pterm.Debug.Printfln("Installing font [%v / %v]: %s", i+1, total, fontURL)

				fontFileName, _ := url.QueryUnescape(filepath.Base(fontURL))
				fontPath := filepath.Join(wrapper.InHome(".local/share/fonts"), fontFileName)

				err := wrapper.Run("wget", "-q", fontURL, "-O", fontPath)
				if err != nil {
					pterm.Fatal.Printfln("Failed to install font %s: %s", fontURL, err)
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
