package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Flashcards",
		Desc: "Studying with flashcards",
		Setup: func() {
			wrapper.FlatpakOnce("io.github.david_swift.Flashcards")
		},
	})
}
