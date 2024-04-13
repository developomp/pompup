package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "FileZilla",
		Desc: "FTP GUI",
		Setup: func() {
			wrapper.FlatpakOnce("org.filezillaproject.Filezilla")
		},
	})
}
