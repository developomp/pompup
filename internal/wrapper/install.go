package wrapper

import (
	"fmt"
	"path/filepath"

	"github.com/pterm/pterm"
)

// Paru installs Arch Linux packages using paru. It can also install
// packages from the AUR.
func Paru(packageName string) error {
	return pacmanLike(packageName, "paru")
}

// ParuOnce acts like Paru but skips when the package is installed already.
func ParuOnce(packageName string) error {
	if IsArchPkgInstalled(packageName) {
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
	if IsArchPkgInstalled(packageName) {
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

// TryDconf tries to load dconf configuration from string then terminates program if it fails.
func TryDconf(data string) {
	err := Dconf(data)
	if err != nil {
		pterm.Fatal.Println("Failed to load config file:", err)
	}
}

// Dconf loads dconf configuration from string
func Dconf(data string) error {
	tmpPath := filepath.Join(TmpDir, "dconf.conf")
	err := WriteFile(tmpPath, []byte(data))
	if err != nil {
		pterm.Fatal.Printfln("Failed to write to %s: %s", tmpPath, err)
	}

	return BashRun(fmt.Sprintf("dconf load / < %s", tmpPath))
}
