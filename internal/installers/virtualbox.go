package installers

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "Virtualbox",
		Desc: "VM stuff",
		Setup: func() {
			wrapper.ParuOnce("virtualbox-host-modules-arch")
			wrapper.ParuOnce("virtualbox")
			wrapper.ParuOnce("virtualbox-ext-oracle")

			wrapper.AddGroup("vboxusers") // allow usage of virtualbox without root perm

			// load vboxdrv, vboxnetadp, and vboxnetflt drivers
			err := wrapper.Run("sudo", "systemctl", "restart", "systemd-modules-load")
			if err != nil {
				pterm.Fatal.Println("Failed to load virtualbox driver modules:", err)
			}

			// set default machine folder
			err = wrapper.Run("VBoxManage", "setproperty", "machinefolder", wrapper.InHome("Dev/OS/VB"))
			if err != nil {
				pterm.Fatal.Println("Failed to set default virtualbox machine folder location:", err)
			}
		},
	})
}
