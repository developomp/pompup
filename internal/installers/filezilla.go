package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "FileZilla",
		Desc: "FTP GUI",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("org.filezillaproject.Filezilla")
		},
	})
}
