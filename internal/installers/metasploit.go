package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "metasploit",
		Desc: "Penetration Testing",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("metasploit")
		},
	})
}
