package install

import (
	"embed"
	"fmt"
	"io"
	"io/fs"

	"github.com/developomp/pompup/internal/helper"
)

//go:embed assets/dconf/extension-*
var _DconfFiles embed.FS
var DconfFiles, _ = fs.Sub(_DconfFiles, "assets/dconf")

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
	// check if package has been installed already
	if helper.BashRun(fmt.Sprintf("flatpak list | grep -w \"%s\"", appID)) == nil {
		return nil // app installed already, nothing to report!
	}

	return helper.Run("flatpak", "install", "-y", "--system", appID)
}

// Dconf loads dconf configuration from internal/workflows/assets/dconf
func Dconf(name string) error {
	file, err := DconfFiles.Open(name)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	tmpFilePath := fmt.Sprintf("/tmp/%s", name)

	err = helper.WriteFile(tmpFilePath, data)
	if err != nil {
		return err
	}

	err = helper.BashRun(fmt.Sprintf("dconf load / <%s", tmpFilePath))
	if err != nil {
		return err
	}

	err = helper.Run("rm", tmpFilePath)
	if err != nil {
		return err
	}

	return nil
}
