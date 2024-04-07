package installers

import (
	_ "embed"
	"os/exec"
	"slices"
	"strings"

	"github.com/developomp/pompup/internal/wrapper"
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
	id      string // GNOME extension ID (numerical)
	archPkg string // Arch Linux package name (including from AUR)
	uuid    string // GNOME extension UUID (XXX@XXX.XX string)
	conf    string // dconf configuration for the extension
}

func init() {
	register(&Installer{
		Name: "GNOME Extensions",
		Desc: "GNOME extensions",
		Tags: []Tag{Gnome},
		Setup: func() {
			wrapper.FlatpakOnce("com.mattjakeman.ExtensionManager") // GNOME extension installer GUI
			wrapper.ParuOnce("gnome-shell-extension-installer")     // GNOME extension installer CLI
			wrapper.ParuOnce("gnome-shell")                         // includes gnome-extensions CLI

			installedExtensions := listInstalledGnomeExtensions()
			extensions := [...]gnomeExtension{
				{"", "gnome-shell-extension-pop-shell-git", "pop-shell@system76.com", _gnomeExtensionPopShellDconf}, // https://aur.archlinux.org/packages/gnome-shell-extension-pop-shell-git
				{"19", "", "user-theme@gnome-shell-extensions.gcampax.github.com", ""},                              // https://extensions.gnome.org/extension/19/user-themes/
				{"36", "", "lockkeys@vaina.lt", _gnomeExtensionLockKeysDconf},                                       // https://extensions.gnome.org/extension/36/lock-keys/
				{"2890", "", "trayIconsReloaded@selfmade.pl", _gnomeExtensionTrayIconsReloadedDconf},                // https://extensions.gnome.org/extension/2890/tray-icons-reloaded/
				{"3193", "", "blur-my-shell@aunetx", _gnomeExtensionBlurMyShellDconf},                               // https://extensions.gnome.org/extension/3193/blur-my-shell/
				{"4158", "", "gnome-ui-tune@itstime.tech", ""},                                                      // https://extensions.gnome.org/extension/4158/gnome-40-ui-improvements/
			}

			pterm.Info.Println("ignore \"Extension X does not exist\" messages below. Extensions are being installed correctly")
			for _, extension := range extensions {
				// install extension
				if len(extension.archPkg) > 0 {
					wrapper.ParuOnce(extension.archPkg)
				} else if len(extension.id) > 0 {
					if !slices.Contains(installedExtensions, extension.uuid) {
						pterm.Debug.Printfln("Installing 'https://extensions.gnome.org/extension/%s' via gnome-shell-extension-installer", extension.id)
						wrapper.Run("gnome-shell-extension-installer", extension.id, "--yes", "--update")
					}
				}

				// apply config
				if len(extension.conf) > 0 {
					err := wrapper.Dconf(extension.conf)

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
			wrapper.TryDconf(_gnomeExtensionEnableDconf)
		},
	})
}

func listInstalledGnomeExtensions() []string {
	out, err := exec.Command("gnome-extensions", "list").Output()
	if err != nil {
		pterm.Fatal.Println("Failed to list installed GNOME extensions:", err)
	}

	return strings.Fields(string(out))
}
