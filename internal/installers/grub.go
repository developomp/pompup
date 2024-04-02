package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GRUB",
		Desc: "Bootloader",
		Tags: []Tag{System},
		Setup: func() {
			if wrapper.IsArchPkgInstalled("grub") {
				return
			}

			wrapper.Paru("grub")

			// https://stackoverflow.com/a/11245501/12979111
			wrapper.Run("sudo", "sed", "-i", "/^GRUB_CMDLINE_LINUX_DEFAULT=/c\\GRUB_CMDLINE_LINUX_DEFAULT=\"loglevel=3 quiet splash nvidia-drm.modeset=1\"", "/etc/default/grub")
			wrapper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT_STYLE=/c\\GRUB_TIMEOUT_STYLE=hidden lol", "/etc/default/grub")
			wrapper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT=/c\\GRUB_TIMEOUT=0", "/etc/default/grub")

			wrapper.Run("sudo", "grub-mkconfig", "-o", "/boot/grub/grub.cfg")
		},
	})
}
