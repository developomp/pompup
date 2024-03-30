package workflows

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Workflow{
		Name: "jdk",
		Desc: "Java Development Kit",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.Paru("jdk-openjdk")
			wrapper.Paru("jdk8-openjdk")
		},
	})
}
