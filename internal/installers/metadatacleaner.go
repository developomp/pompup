package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Metadata Cleaner",
		Desc: "Metadata Cleaner",
		Setup: func() {
			wrapper.FlatpakOnce("fr.romainvigier.MetadataCleaner")
		},
	})
}
