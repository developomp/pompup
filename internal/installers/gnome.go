package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/pipewire/pipewire.conf.d/99-low-latency.conf
var _pipewireLowLatencyConfig string

//go:embed assets/dconf/gnome-desktop.conf
var _gnomeDesktopDconf string

//go:embed assets/dconf/gnome-keybindings.conf
var _gnomeKeybindings string

//go:embed assets/dconf/gnome-input.conf
var _gnomeInputDconf string

//go:embed assets/dconf/gnome-mutter.conf
var _gnomeMutterDconf string

//go:embed assets/dconf/gnome-settings-daemon.conf
var _gnomeSettingsDaemon string

//go:embed assets/dconf/extension-user-themes.conf
var _gnomeExtensionUserThemesDconf string

func init() {
	register(&Installer{
		Name: "GNOME",
		Desc: "minimal, usable GNOME Desktop Environment (DE)",
		Setup: func() {
			// core
			wrapper.ParuOnce("gdm")
			// gdm
			// └── gnome-shell
			//     └── mutter
			wrapper.ParuOnce("gnome-control-center") // installs gnome-settings-daemon as dependency
			// gnome-control-center
			// └── gnome-settings-daemon
			wrapper.ParuOnce("xdg-desktop-portal-gnome")
			// xdg-desktop-portal-gnome
			// └── gnome-keyring

			wrapper.ParuOnce("touchegg") // for wayland touch gestures
			wrapper.ParuOnce("ibus")
			wrapper.ParuOnce("ibus-hangul") // Korean input
			wrapper.ParuOnce("ibus-anthy")  // Japanese input
			// todo: AUR not building as of the moment, re-enable later
			// wrapper.ParuOnce("ibus-pinyin") // Chinese input

			// themes
			wrapper.ParuOnce("pop-gtk-theme")      // gtk theme
			wrapper.ParuOnce("papirus-icon-theme") // icon theme
			wrapper.ParuOnce("xcursor-breeze")     // cursor theme

			configurePipewire()
			wrapper.TryDconf(_gnomeDesktopDconf)
			wrapper.TryDconf(_gnomeKeybindings)
			wrapper.TryDconf(_gnomeInputDconf)
			wrapper.TryDconf(_gnomeMutterDconf)
			wrapper.TryDconf(_gnomeSettingsDaemon)
			wrapper.TryDconf(_gnomeExtensionUserThemesDconf)

			err := wrapper.Run("sudo", "systemctl", "enable", "gdm")
			if err != nil {
				pterm.Fatal.Println("Failed to enable gdm:", err)
			}

			// set gdm login screen mouse settings
			err = wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.peripherals.mouse", "accel-profile", "flat")
			if err != nil {
				pterm.Fatal.Println("Failed to enable gdm:", err)
			}
			err = wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.peripherals.mouse", "speed", "0.0")
			if err != nil {
				pterm.Fatal.Println("Failed to enable gdm:", err)
			}
			err = wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.interface", "cursor-theme", "Breeze")
			if err != nil {
				pterm.Fatal.Println("Failed to enable gdm:", err)
			}
		},
	})
}

// configurePipewire configures pipewire to be optimized for low latency while
// sacrificing CPU performance. This function requires pipewire to be already
// installed (pipewire is part of gdm's dependency chain: gdm > gnome-shell > mutter > pipewire).
// more info: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire
// more info about split conf: https://gitlab.freedesktop.org/pipewire/pipewire/-/wikis/Config-PipeWire#split-file-configuration
func configurePipewire() {
	wrapper.SudoWriteFile("/etc/pipewire/pipewire.conf.d/99-low-latency.conf", _pipewireLowLatencyConfig)

	// restart pipewire for the config to take effect
	err := wrapper.Run("systemctl", "--user", "restart", "pipewire.service")
	if err != nil {
		pterm.Fatal.Println("Failed to restart pipewire service:", err)
	}
}
