package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "metasploit",
		Desc: "Penetration Testing",
		Setup: func() {
			wrapper.ParuOnce("metasploit")
		},
	})
}
