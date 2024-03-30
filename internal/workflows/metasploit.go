package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "metasploit",
		Desc: "Penetration Testing",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			wrapper.Paru("metasploit")
		},
	})
}
