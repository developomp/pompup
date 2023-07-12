package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "jdk",
		Desc: "Java Development Kit",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			install.Paru("jdk-openjdk")
			install.Paru("jdk8-openjdk")
		},
	})
}
