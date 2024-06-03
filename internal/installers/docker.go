package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "docker",
		Desc: "not a VM",
		Setup: func() {
			wrapper.ParuOnce("docker")
			wrapper.ParuOnce("docker-compose")
			wrapper.AddGroup("docker") // allow use of docker CLI without sudo
			wrapper.SystemctlEnable("docker", wrapper.EnableNow)
		},
	})
}
