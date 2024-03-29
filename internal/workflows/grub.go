package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "GRUB",
		Desc: "Bootloader",
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("grub")

			// https://stackoverflow.com/a/11245501/12979111
			helper.Run("sudo", "sed", "-i", "/^GRUB_CMDLINE_LINUX_DEFAULT=/c\\GRUB_CMDLINE_LINUX_DEFAULT=\"loglevel=3 quiet splash nvidia-drm.modeset=1\"", "/etc/default/grub")
			helper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT_STYLE=/c\\GRUB_TIMEOUT_STYLE=hidden lol", "/etc/default/grub")
			helper.Run("sudo", "sed", "-i", "/^GRUB_TIMEOUT=/c\\GRUB_TIMEOUT=0", "/etc/default/grub")

			helper.Run("sudo", "grub-mkconfig", "-o", "/boot/grub/grub.cfg")
		},
	})
}
