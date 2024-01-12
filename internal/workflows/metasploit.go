package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "metasploit",
		Desc: "Penetration Testing",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("metasploit")
		},
	})
}
