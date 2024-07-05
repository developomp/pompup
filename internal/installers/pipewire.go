package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/pipewire/pipewire.conf.d/99-low-latency.conf
var _pipewireLowLatencyConfig string

//go:embed assets/etc/wireplumber/wireplumber.conf.d/99-low-latency.conf
var _wireplumberConfig string

func init() {
	register(&Installer{
		Name:     "Pipewire",
		Desc:     "audio and video for Linux",
		Priority: -1,
		Setup: func() {
			// - configuration optimized for low latency rhythm game
			// - 48kHz sample rate only
			//
			// Audio latency learning resources:
			//   - https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/FAQ#pipewire-buffering-explained
			//   - https://pipewire.pages.freedesktop.org/wireplumber/daemon/configuration/alsa.html#alsa-buffer-properties
			//   - https://www.reddit.com/r/linuxaudio/comments/v66mj6/comment/ibjupwb
			//   - https://learn.microsoft.com/en-us/windows-hardware/drivers/audio/low-latency-audio

			wrapper.ParuOnce("pipewire")
			wrapper.ParuOnce("pipewire-alsa")
			wrapper.ParuOnce("pipewire-jack")
			wrapper.ParuOnce("pipewire-pulse")
			wrapper.ParuOnce("wireplumber")

			wrapper.SudoWriteFile("/etc/pipewire/pipewire.conf.d/99-low-latency.conf", _pipewireLowLatencyConfig)
			wrapper.SudoWriteFile("/etc/wireplumber/wireplumber.conf.d/99-low-latency.conf", _wireplumberConfig)

			// restart pipewire for the config to take effect
			err := wrapper.Run("systemctl", "--user", "restart", "pipewire.service")
			if err != nil {
				pterm.Fatal.Println("Failed to restart pipewire service:", err)
			}
		},
	})
}
