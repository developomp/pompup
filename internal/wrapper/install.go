package wrapper

import (
	"fmt"

	"github.com/pterm/pterm"
)

// Paru installs Arch Linux packages using paru. It can also install
// packages from the AUR.
func Paru(packageName string) error {
	return pacmanLike(packageName, "paru")
}

// ParuOnce acts like Paru but skips when the package is installed already.
func ParuOnce(packageName string) error {
	if IsArchPkgInstalled("paru", packageName) {
		return nil
	}

	return pacmanLike(packageName, "paru")
}

// Pacman installs Arch Linux packages using pacman. It can not install packages
// from the AUR.
func Pacman(packageName string) error {
	return pacmanLike(packageName, "pacman")
}

// PacmanOnce acts like Pacman but skips when the package is installed already.
func PacmanOnce(packageName string) error {
	if IsArchPkgInstalled("pacman", packageName) {
		return nil
	}

	return pacmanLike(packageName, "pacman")
}

// Flatpak installs flatpak packages.
func Flatpak(appID string) error {
	// check if package has been installed already
	if BashRun(fmt.Sprintf("flatpak list | grep -w \"%s\"", appID)) == nil {
		return nil // app installed already, nothing to report!
	}

	pterm.Debug.Printfln("Installing '%s' via flatpak", appID)
	return Run("flatpak", "install", "-y", "--system", appID)
}

// FlatpakOnce acts like Flatpak but skips when the package is installed already.
func FlatpakOnce(packageName string) error {
	if IsFlatpakInstalled(packageName) {
		return nil
	}

	return Flatpak(packageName)
}

// Dconf loads dconf configuration from string
func Dconf(data string) error {
	return BashRun(fmt.Sprintf("dconf load / << EOF\n%s\nEOF", data))
}
