package wrapper

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

// pacmanLike is a template for pacman and pacman-like software
// (pacman, yay, paru, etc...). It is only here to reduce code duplication
func pacmanLike(packageName string, installer string) error {
	pterm.Debug.Printfln("Installing '%s' via %s", packageName, installer)

	// Skip installation if the package is installed already.
	if IsArchPkgInstalled(packageName) {
		return nil
	}

	var cmd *exec.Cmd
	if installer == "pacman" {
		// pacman needs sudo to install packages
		cmd = exec.Command("bash", "-c", fmt.Sprintf("yes | sudo %s -S --needed %s", installer, packageName))
	} else {
		cmd = exec.Command("bash", "-c", fmt.Sprintf("yes |      %s -S --needed %s", installer, packageName))
	}
	cmd.Stdout = os.Stdout // show stdout
	cmd.Stderr = os.Stderr // show stderr

	return cmd.Run()
}
