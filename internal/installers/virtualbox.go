package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Virtualbox",
		Desc: "VM stuff",
		Tags: []Tag{Dev, Gui},
		Setup: func() {
			wrapper.ParuOnce("virtualbox")
			wrapper.ParuOnce("virtualbox-host-modules-arch")
			wrapper.ParuOnce("virtualbox-ext-oracle")

			wrapper.AddGroup("vboxusers") // allow usage of virtualbox without root perm

			// load vboxdrv, vboxnetadp, and vboxnetflt drivers
			err := wrapper.Run("sudo", "systemctl", "restart", "systemd-modules-load")
			if err != nil {
				pterm.Fatal.Println("Failed to load virtualbox driver modules:", err)
			}
		},
		Reminders: []string{
			"Re-login to use virtualbox without sudo",
		},
	})
}
