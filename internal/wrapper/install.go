package wrapper

import (
	"fmt"
)

// Paru installs Arch Linux packages using paru. It can also install
// packages from the AUR.
func Paru(packageName string) error {
	return pacmanLike(packageName, "paru")
}

// Pacman installs Arch Linux packages using pacman. It can not install packages
// from the AUR.
func Pacman(packageName string) error {
	return pacmanLike(packageName, "pacman")
}

// Flatpak installs flatpak packages.
func Flatpak(appID string) error {
	// check if package has been installed already
	if BashRun(fmt.Sprintf("flatpak list | grep -w \"%s\"", appID)) == nil {
		return nil // app installed already, nothing to report!
	}

	return Run("flatpak", "install", "-y", "--system", appID)
}

// Dconf loads dconf configuration from string
func Dconf(data string) error {
	return BashRun(fmt.Sprintf("dconf load / << EOF\n%s\nEOF", data))
}
