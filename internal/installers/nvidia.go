package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func init() {
	register(&Installer{
		Name:     "Nvidia",
		Desc:     "🖕",
		Priority: -1,
		Setup: func() {
			// https://wiki.archlinux.org/title/NVIDIA
			// https://wiki.archlinux.org/title/OpenGL
			// https://wiki.archlinux.org/title/Vulkan
			// https://wiki.archlinux.org/title/Hardware_video_acceleration
			// related: environment.go, mkinitcpio.go

			wrapper.ParuOnce("nvidia")
			// nvidia
			// └── nvidia-utils <- automatically blacklists the nouveau module
			wrapper.ParuOnce("cuda")
			// cuda
			// └── opencl-nvidia
			wrapper.ParuOnce("nvidia-settings")

			wrapper.ParuOnce("mesa")

			// https://wiki.archlinux.org/title/GDM#Wayland_and_the_proprietary_NVIDIA_driver
			if !wrapper.PathExists("/etc/udev/rules.d/61-gdm.rules") {
				err := wrapper.Run("sudo", "ln", "-s", "/dev/null", "/etc/udev/rules.d/61-gdm.rules")
				if err != nil {
					pterm.Fatal.Println("Failed to run mkinitcpio:", err)
				}
			}

			// https://wiki.archlinux.org/title/NVIDIA/Tips_and_tricks#Preserve_video_memory_after_suspend
			wrapper.SystemctlEnable("nvidia-suspend", wrapper.EnableOnly)
			wrapper.SystemctlEnable("nvidia-hibernate", wrapper.EnableOnly)
			wrapper.SystemctlEnable("nvidia-resume", wrapper.EnableOnly)
		},
	})
}
