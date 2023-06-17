package install

import (
	"os/exec"
)

// Deps installs dependencies required by the application.
// It only needs to be called once during the initialization process.
func Deps() {
	Pacman("trash-cli")
	Pacman("flatpak")

	installParu()
}

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
	return exec.Command("flatpak", "install", "-y", "--system", appID).Run()
}

// Dconf loads dconf configuration from file
func Dconf(filePath string) error {
	return exec.Command("dconf", "load", "/", "<", filePath).Run()
}