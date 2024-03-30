package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "FileZilla",
		Desc: "FTP GUI",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.Flatpak("org.filezillaproject.Filezilla")
		},
	})
}
