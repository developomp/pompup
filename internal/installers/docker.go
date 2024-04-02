package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "docker",
		Desc: "not a VM",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			wrapper.ParuOnce("docker")

			// allow use of docker CLI without sudo
			err := wrapper.AddGroup("docker")
			if err != nil {
				pterm.Fatal.Println("Failed to add user to docker group:", err)
			}

			err = wrapper.Run("systemctl", "is-active", "docker")
			if err != nil {
				err = wrapper.Run("sudo", "systemctl", "--now", "enable", "docker")
				if err != nil {
					pterm.Fatal.Println("Failed to start docker service:", err)
				}
			}
		},
	})
}
