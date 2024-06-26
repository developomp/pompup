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
func ParuOnce(packageName string) {
	if IsArchPkgInstalled(packageName) {
		return
	}

	err := pacmanLike(packageName, "paru")
	if err != nil {
		pterm.Fatal.Printfln("Failed to install %s via paru: %s", packageName, err)
	}
}

// Pacman installs Arch Linux packages using pacman. It can not install packages
// from the AUR.
func Pacman(packageName string) error {
	return pacmanLike(packageName, "pacman")
}

// PacmanOnce acts like Pacman but skips when the package is installed already.
func PacmanOnce(packageName string) {
	if IsArchPkgInstalled(packageName) {
		return
	}

	err := pacmanLike(packageName, "pacman")
	if err != nil {
		pterm.Fatal.Printfln("Failed to install %s via pacman: %s", packageName, err)
	}
}

// Flatpak installs flatpak packages.
func Flatpak(appID string) error {
	pterm.Debug.Printfln("Installing '%s' via flatpak", appID)
	return Run("flatpak", "install", "-y", "--system", appID)
}

// FlatpakOnce acts like Flatpak but skips when the package is installed already.
func FlatpakOnce(packageName string) {
	if IsFlatpakInstalled(packageName) {
		return
	}

	err := Flatpak(packageName)
	if err != nil {
		pterm.Fatal.Printfln("Failed to install %s via flatpak: %s", packageName, err)
	}
}

// Nix installs nix packages.
func Nix(packageName string) error {
	pterm.Debug.Printfln("Installing '%s' via nix", packageName)
	return Run("nix", "profile", "install", packageName, "--impure")
}

// NixOnce acts like Nix but skips when the package is installed already.
func NixOnce(packageName string) {
	if IsNixInstalled(packageName) {
		return
	}

	err := Nix(packageName)
	if err != nil {
		pterm.Fatal.Printfln("Failed to install %s via nix: %s", packageName, err)
	}
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
	WriteFile(tmpPath, []byte(data))

	return BashRun(fmt.Sprintf("dconf load / < %s", tmpPath))
}
