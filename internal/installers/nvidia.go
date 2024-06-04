package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/mkinitcpio.conf
var _mkinitcpio string

// https://wiki.archlinux.org/title/NVIDIA
// https://wiki.archlinux.org/title/OpenGL
// https://wiki.archlinux.org/title/Vulkan
// https://wiki.archlinux.org/title/Hardware_video_acceleration
// related: grub.go
func init() {
	register(&Installer{
		Name:     "Nvidia",
		Desc:     "🖕",
		Priority: -1,
		Setup: func() {
			wrapper.ParuOnce("nvidia")
			// nvidia
			// └── nvidia-utils
			wrapper.ParuOnce("cuda")
			// cuda
			// └── opencl-nvidia
			wrapper.ParuOnce("nvidia-settings")

			const configPath = "/etc/mkinitcpio.conf"
			if !wrapper.IsFileUpdated(configPath, _mkinitcpio) {
				wrapper.SudoWriteFile(configPath, _mkinitcpio)

				err := wrapper.Run("sudo", "mkinitcpio", "-P")
				if err != nil {
					pterm.Fatal.Println("Failed to run mkinitcpio:", err)
				}
			}

			// https://wiki.archlinux.org/title/GDM#Wayland_and_the_proprietary_NVIDIA_driver
			err := wrapper.Run("sudo", "ln", "-s", "/dev/null", "/etc/udev/rules.d/61-gdm.rules")
			if err != nil {
				pterm.Fatal.Println("Failed to run mkinitcpio:", err)
			}

			// https://wiki.archlinux.org/title/NVIDIA/Tips_and_tricks#Preserve_video_memory_after_suspend
			wrapper.SystemctlEnable("nvidia-suspend", wrapper.EnableOnly)
			wrapper.SystemctlEnable("nvidia-hibernate", wrapper.EnableOnly)
			wrapper.SystemctlEnable("nvidia-resume", wrapper.EnableOnly)
		},
	})
}
