package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/mkinitcpio.conf
var _mkinitcpio string

//go:embed assets/etc/modprobe.d/nvidia.conf
var _nvidiaModprobe string

func init() {
	register(&Installer{
		Name:     "Mkinitcpio",
		Desc:     "RAMdisk stuff",
		Priority: 1,
		Setup: func() {
			// mkinitcpio is a dependency of linux. No need to install them.

			wrapper.SudoWriteFile("/etc/modprobe.d/nvidia.conf", _nvidiaModprobe)

			const configPath = "/etc/mkinitcpio.conf"
			if !wrapper.IsFileUpdated(configPath, _mkinitcpio) {
				wrapper.SudoWriteFile(configPath, _mkinitcpio)

				err := wrapper.Run("sudo", "mkinitcpio", "-P")
				if err != nil {
					pterm.Fatal.Println("Failed to run mkinitcpio:", err)
				}
			}
		},
	})
}
