package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Flashcards",
		Desc: "Studying with flashcards",
		Tags: []Tag{Gui},
		Setup: func() {
			wrapper.FlatpakOnce("io.github.david_swift.Flashcards")
		},
	})
}
