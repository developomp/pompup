package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/default/grub
var _grubConfig string

func init() {
	register(&Installer{
		Name: "GRUB",
		Desc: "Bootloader",
		Setup: func() {
			wrapper.ParuOnce("grub")

			const configPath = "/etc/default/grub"
			if !wrapper.IsFileUpdated(configPath, _grubConfig) {
				wrapper.SudoWriteFile(configPath, _grubConfig)

				err := wrapper.Run("sudo", "grub-mkconfig", "-o", "/boot/grub/grub.cfg")
				if err != nil {
					pterm.Fatal.Println("Failed to generate /boot/grub/grub.cfg:", err)
				}
			}
		},
	})
}
