package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/dconf/extension-enable.conf
var _gnomeExtensionEnableDconf string

//go:embed assets/dconf/extension-pop-shell.conf
var _gnomeExtensionPopShellDconf string

//go:embed assets/dconf/extension-lockkeys.conf
var _gnomeExtensionLockKeysDconf string

//go:embed assets/dconf/extension-trayIconsReloaded.conf
var _gnomeExtensionTrayIconsReloadedDconf string

//go:embed assets/dconf/extension-blur-my-shell.conf
var _gnomeExtensionBlurMyShellDconf string

type gnomeExtension = struct {
	id      string // GNOME extension ID
	archPkg string // Arch Linux package name (including from AUR)
	conf    string // dconf configuration for the extension
}

func init() {
	register(&Workflow{
		Name: "GNOME Extensions",
		Desc: "GNOME extensions",
		Tags: []Tag{Gnome},
		Setup: func() {
			install.Flatpak("com.mattjakeman.ExtensionManager") // GNOME extension installer GUI
			install.Paru("gnome-shell-extension-installer")     // GNOME extension installer CLI

			extensions := [...]gnomeExtension{
				{"", "gnome-shell-extension-pop-shell-git", _gnomeExtensionPopShellDconf}, // https://aur.archlinux.org/packages/gnome-shell-extension-pop-shell-git
				{"19", "", ""},                                      // https://extensions.gnome.org/extension/19/user-themes/
				{"36", "", _gnomeExtensionLockKeysDconf},            // https://extensions.gnome.org/extension/36/lock-keys/
				{"2890", "", _gnomeExtensionTrayIconsReloadedDconf}, // https://extensions.gnome.org/extension/2890/tray-icons-reloaded/
				{"3193", "", _gnomeExtensionBlurMyShellDconf},       // https://extensions.gnome.org/extension/3193/blur-my-shell/
				{"4158", "", ""},                                    // https://extensions.gnome.org/extension/4158/gnome-40-ui-improvements/
			}

			pterm.Info.Println("ignore \"Extension X does not exist\" messages below. Extensions are being installed correctly")
			for _, extension := range extensions {
				// install extension
				if len(extension.archPkg) > 0 {
					install.Paru(extension.archPkg)
				} else if len(extension.id) > 0 {
					pterm.Debug.Printfln("Installing 'https://extensions.gnome.org/extension/%s' via gnome-shell-extension-installer", extension.id)
					helper.Run("gnome-shell-extension-installer", extension.id, "--yes", "--update")
				}

				// apply config
				if len(extension.conf) > 0 {
					err := install.Dconf(extension.conf)

					if err != nil {
						if len(extension.archPkg) > 0 {
							pterm.Fatal.Printfln("Failed to load dconf config for extension '%s' (arch package): %s", extension.archPkg, err)
						} else if len(extension.id) > 0 {
							pterm.Fatal.Printfln("Failed to load dconf config for 'https://extensions.gnome.org/extension/%s': %s", extension.id, err)
						}
					}
				}
			}

			// enable extensions
			install.Dconf(_gnomeExtensionEnableDconf)
		},
		Reminders: []string{
			"Log-out and Log-in to GNOME",
		},
	})
}
