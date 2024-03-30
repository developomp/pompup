package wrapper

import (
	_ "embed"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
)

// pacmanLike is a template for pacman and pacman-like software
// (pacman, yay, paru, etc...). It is only here to reduce code duplication
func pacmanLike(packageName string, installer string) error {
	pterm.Debug.Printfln("Installing '%s' via %s", packageName, installer)

	// Skip installation if the package is installed already.
	if IsArchPkgInstalled(installer, packageName) {
		return nil
	}

	var cmd *exec.Cmd
	if installer == "pacman" {
		// pacman needs sudo to run
		cmd = exec.Command("sudo", installer, "-S", "--noconfirm", "--needed", packageName)
	} else {
		cmd = exec.Command( /*  */ installer, "-S", "--noconfirm", "--needed", packageName)
	}
	cmd.Stdout = os.Stdout // show stdout
	cmd.Stderr = os.Stderr // show stderr

	return cmd.Run()
}
