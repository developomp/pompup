package workflows

import (
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

type gnomeExtension = struct {
	id      string // GNOME extension ID
	archPkg string // Arch Linux package name (including from AUR)

	confName string // dconf name of the extension
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
				{"", "gnome-shell-extension-pop-shell-git", "extension-pop-shell.conf"}, // https://aur.archlinux.org/packages/gnome-shell-extension-pop-shell-git
				{"36", "", "extension-lockkeys.conf"},                                   // https://extensions.gnome.org/extension/36/lock-keys/
				{"2890", "", "extension-trayIconsReloaded.conf"},                        // https://extensions.gnome.org/extension/2890/tray-icons-reloaded/
				{"3193", "", "extension-blur-my-shell.conf"},                            // https://extensions.gnome.org/extension/3193/blur-my-shell/
				{"4158", "", ""}, // https://extensions.gnome.org/extension/4158/gnome-40-ui-improvements/
			}

			pterm.Info.Println("ignore \"Extension X does not exist\" messages below. Extensions are being installed correctly")
			for _, extension := range extensions {
				// log("installing: https://extensions.gnome.org/extension/$1")
				if len(extension.archPkg) > 0 {
					install.Paru(extension.archPkg)
				} else if len(extension.id) > 0 {
					pterm.Debug.Printfln("Installing 'https://extensions.gnome.org/extension/%s' via gnome-shell-extension-installer", extension.id)
					helper.Run("gnome-shell-extension-installer", extension.id, "--yes", "--update")
				}

				if len(extension.confName) > 0 {
					err := install.Dconf(extension.confName)

					if err != nil {
						pterm.Fatal.Printfln("Failed to load dconf config '%s':%s", extension.confName, err)
					}
				}
			}

			// enable extensions
			install.Dconf("extension-enable.conf")
		},
		Reminders: []string{
			"Log-out and Log-in to GNOME",
		},
	})
}
