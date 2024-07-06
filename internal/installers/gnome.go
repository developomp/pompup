package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/dconf/gnome-desktop.conf
var _gnomeDesktopDconf string

//go:embed assets/dconf/gnome-keybindings.conf
var _gnomeKeybindings string

//go:embed assets/dconf/gnome-input.conf
var _gnomeInputDconf string

//go:embed assets/dconf/gnome-mutter.conf
var _gnomeMutterDconf string

//go:embed assets/dconf/gnome-power.conf
var _gnomePowerDconf string

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

			wrapper.TryDconf(_gnomeExtensionUserThemesDconf)
			wrapper.TryDconf(_gnomeDesktopDconf)
			wrapper.TryDconf(_gnomeKeybindings)
			wrapper.TryDconf(_gnomeInputDconf)
			wrapper.TryDconf(_gnomeMutterDconf)
			wrapper.TryDconf(_gnomePowerDconf)
			wrapper.TryDconf(_gnomeSettingsDaemon)

			wrapper.SystemctlEnable("gdm", wrapper.EnableOnly)
			wrapper.SystemctlEnable("touchegg", wrapper.EnableNow)

			// set gdm login screen mouse settings
			err := wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.peripherals.mouse", "accel-profile", "flat")
			if err != nil {
				pterm.Fatal.Println("Failed to set gdm login screen mouse acceleration:", err)
			}
			err = wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.peripherals.mouse", "speed", "0.0")
			if err != nil {
				pterm.Fatal.Println("Failed to set gdm login screen mouse speed:", err)
			}
			err = wrapper.Run("sudo", "-u", "gdm", "dbus-launch", "gsettings", "set", "org.gnome.desktop.interface", "cursor-theme", "Breeze")
			if err != nil {
				pterm.Fatal.Println("Failed to set gdm login screen mouse cursor theme:", err)
			}
		},
	})
}
