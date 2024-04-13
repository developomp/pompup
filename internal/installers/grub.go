package installers

import (
	"fmt"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name: "GRUB",
		Desc: "Bootloader",
		Setup: func() {
			wrapper.ParuOnce("grub")

			changed := false
			const filePath = "/etc/default/grub"

			// The sed commands below are explained here: https://stackoverflow.com/a/11245501/12979111

			if !isLineInFile("GRUB_CMDLINE_LINUX_DEFAULT=\"loglevel=3 quiet splash nvidia-drm.modeset=1\"", filePath) {
				changed = true
				wrapper.Run("sudo", "sed", "-i", "/^GRUB_CMDLINE_LINUX_DEFAULT=/c\\GRUB_CMDLINE_LINUX_DEFAULT=\"loglevel=3 quiet splash nvidia-drm.modeset=1\"", filePath)
			}

			if !isLineInFile("GRUB_TIMEOUT_STYLE=hidden", filePath) {
				changed = true
				wrapper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT_STYLE=/c\\GRUB_TIMEOUT_STYLE=hidden", filePath)
			}

			if !isLineInFile("GRUB_TIMEOUT=0", filePath) {
				changed = true
				wrapper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT=/c\\GRUB_TIMEOUT=0", filePath)
			}

			if changed {
				err := wrapper.Run("sudo", "grub-mkconfig", "-o", "/boot/grub/grub.cfg")
				if err != nil {
					pterm.Fatal.Println("Failed to generate /boot/grub/grub.cfg:", err)
				}
			}
		},
	})
}

func isLineInFile(line string, filePath string) bool {
	// ^: beginning of line
	// $: end of line
	err := wrapper.Run("grep", fmt.Sprintf("^%s$", line), filePath)
	return err == nil
}
