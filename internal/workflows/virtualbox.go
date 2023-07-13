package workflows

import (
	"os/exec"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "Virtualbox",
		Desc: "VM stuff",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			install.Paru("virtualbox")
			install.Paru("virtualbox-host-modules-arch")
			install.Paru("virtualbox-ext-oracle")

			username := getUserName()

			// allow usage of virtualbox without root perm
			err := exec.Command("sudo", "gpasswd", "-a", username, "vboxusers").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to add", username, "to vboxusers group:", err)
			}

			// load vboxdrv, vboxnetadp, and vboxnetflt drivers
			err = exec.Command("sudo", "systemctl", "restart", "systemd-modules-load").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to load virtualbox driver modules:", err)
			}
		},
		Reminders: []string{
			"Re-login to use virtualbox without sudo",
		},
	})
}
