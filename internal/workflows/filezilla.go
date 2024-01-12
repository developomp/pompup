package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "FileZilla",
		Desc: "FTP GUI",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Flatpak("org.filezillaproject.Filezilla")
		},
	})
}
