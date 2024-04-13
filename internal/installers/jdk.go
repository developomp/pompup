package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "jdk",
		Desc: "Java Development Kit",
		Setup: func() {
			wrapper.ParuOnce("jdk-openjdk")
			wrapper.ParuOnce("jdk8-openjdk")
		},
	})
}
