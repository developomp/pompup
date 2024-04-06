package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "docker",
		Desc: "not a VM",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.ParuOnce("docker")
			wrapper.AddGroup("docker") // allow use of docker CLI without sudo
			wrapper.SystemctlEnable("docker")
		},
	})
}
