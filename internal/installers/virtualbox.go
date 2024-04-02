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

			username := wrapper.GetUserName()

			// allow usage of virtualbox without root perm
			err := wrapper.Run("sudo", "gpasswd", "-a", username, "vboxusers")
			if err != nil {
				pterm.Fatal.Println("Failed to add", username, "to vboxusers group:", err)
			}

			// load vboxdrv, vboxnetadp, and vboxnetflt drivers
			err = wrapper.Run("sudo", "systemctl", "restart", "systemd-modules-load")
			if err != nil {
				pterm.Fatal.Println("Failed to load virtualbox driver modules:", err)
			}
		},
		Reminders: []string{
			"Re-login to use virtualbox without sudo",
		},
	})
}
