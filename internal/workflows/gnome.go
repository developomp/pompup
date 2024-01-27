package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/pipewire/pipewire.conf.d/99-low-latency.conf
var _pipewireLowLatencyConfig string

//go:embed assets/dconf/gnome-desktop.conf
var _gnomeDesktopDconf string

//go:embed assets/dconf/gnome-keybindings.conf
var _gnomeKeybindings string

//go:embed assets/dconf/gnome-mutter.conf
var _gnomeMutterDconf string

//go:embed assets/dconf/gnome-settings-daemon.conf
var _gnomeSettingsDaemon string

//go:embed assets/dconf/extension-user-themes.conf
var _gnomeExtensionUserThemesDconf string

func init() {
	register(&Workflow{
		Name: "GNOME",
		Desc: "minimal, usable GNOME Desktop Environment (DE)",
		Tags: []Tag{System},
		Setup: func() {
			// core
			install.Paru("gdm")
			install.Paru("gnome-keyring")
			install.Paru("gnome-control-center") // installs gnome-settings-daemon as dependency
			install.Paru("xdg-desktop-portal-gnome")
			install.Paru("touchegg") // for wayland touch gestures

			// themes
			install.Paru("pop-gtk-theme")      // gtk theme
			install.Paru("papirus-icon-theme") // icon theme
			install.Paru("xcursor-breeze")     // cursor theme

			configurePipewire()
			install.Dconf(_gnomeDesktopDconf)
			install.Dconf(_gnomeKeybindings)
			install.Dconf(_gnomeMutterDconf)
			install.Dconf(_gnomeSettingsDaemon)
			install.Dconf(_gnomeExtensionUserThemesDconf)
		},
	})
}

// configurePipewire configures pipewire to be optimized for low latency while
// sacrificing CPU performance. This function requires pipewire to be already
// installed (pipewire is part of gdm's dependency chain: gdm > gnome-shell > mutter > pipewire).
// more info: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire
// more info about split conf: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire#split-file-configuration
func configurePipewire() {
	err := helper.SudoWriteFile("/etc/pipewire/pipewire.conf.d/99-low-latency.conf", _pipewireLowLatencyConfig)
	if err != nil {
		pterm.Fatal.Println("Failed to configure pipewire")
	}

	// restart pipewire for the config to take effect
	err = helper.Run("systemctl", "--user", "restart", "pipewire.service")
}
