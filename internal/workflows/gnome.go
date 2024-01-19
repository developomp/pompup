package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/pipewire/pipewire.conf.d/99-low-latency.conf
var pipewireLowLatencyConfig string

func init() {
	register(&Workflow{
		Name: "GNOME",
		Desc: "GNOME desktop environment",
		Tags: []Tag{System},
		Setup: func() {
			install.Paru("gdm")
			install.Paru("gnome-keyring")

			configurePipewire()
		},
	})
}

// configurePipewire configures pipewire to be optimized for low latency while
// sacrificing CPU performance. This function requires pipewire to be already
// installed (pipewire is part of gdm's dependency chain: gdm > gnome-shell > mutter > pipewire).
// more info: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire
// more info about split conf: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire#split-file-configuration
func configurePipewire() {
	err := helper.SudoWriteFile("/etc/pipewire/pipewire.conf.d/99-low-latency.conf", pipewireLowLatencyConfig)
	if err != nil {
		pterm.Fatal.Println("Failed to configure pipewire")
	}

	// restart pipewire for the config to take effect
	err = helper.Run("systemctl", "--user", "restart", "pipewire.service")
}
