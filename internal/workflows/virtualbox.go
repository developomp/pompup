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

			// allow usage of virtualbox without root perm
			err := exec.Command("sudo", "systemctl", "gpasswd", "-a", getUserName(), "vboxusers").Run()
			if err != nil {
				pterm.Fatal.Println(":", err)
			}

			// load vboxdrv, vboxnetadp, and vboxnetflt drivers
			err = exec.Command("sudo", "systemctl", "restart", "systemd-modules-load").Run()
			if err != nil {
				pterm.Fatal.Println(":", err)
			}
		}, PostSetup: func() {
			pterm.Info.Println("Re-login to use virtualbox without sudo")
		},
	})
}
