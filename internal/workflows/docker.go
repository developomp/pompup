package workflows

import (
	"os/exec"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "docker",
		Desc: "not a VM",
		Tags: []Tag{Cli, Dev},
		Setup: func() {
			install.Paru("docker")

			// add current user to the "docker" group
			// this allows us to use docker commands without sudo
			username := helper.GetUserName() // assumes there's no space in the username
			err := exec.Command("sudo", "usermod", "-aG", "docker", username).Run()
			if err != nil {
				pterm.Fatal.Println("Failed to add", username, "to docker group:", err)
			}

			err = exec.Command("sudo", "systemctl", "--now", "enable", "docker").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to start docker service:", err)
			}
		},
	})
}
