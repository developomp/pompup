package installers

import (
	"os/exec"
	"strings"

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

			out, err := exec.Command("groups").Output()
			if err != nil {
				pterm.Fatal.Println("Failed to get group names:", err)
			}
			if !strings.Contains(string(out), "docker") {
				// add current user to the "docker" group
				// this allows us to use docker commands without sudo
				username := wrapper.GetUserName() // assumes there's no space in the username
				err = wrapper.Run("sudo", "usermod", "-aG", "docker", username)
				if err != nil {
					pterm.Fatal.Println("Failed to add", username, "to docker group:", err)
				}
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
