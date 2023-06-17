package install

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/check"
	"github.com/developomp/pompup/internal/constants"
	"github.com/pterm/pterm"
)

// pacmanLike is a template for pacman and pacman-like software
// (pacman, yay, paru, etc...). It is only here to reduce code duplication
func pacmanLike(packageName string, installer string) error {
	pterm.Info.Printfln("Installing '%s' via %s", packageName, installer)

	// Skip installation if the package is installed already.
	// Usage of bash is a hack that allows me to use pipe.
	if err := exec.Command("bash", "-c", fmt.Sprintf("%s -Q | grep %v", installer, packageName)).Run(); err == nil {
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

// installParu installs paru if it's not installed already.
func installParu() {
	if check.IsInstalled("paru") {
		return
	}

	Pacman("git")
	Pacman("base-devel")

	var cmd *exec.Cmd

	cmd = exec.Command("git", "clone", "https://aur.archlinux.org/paru.git")
	cmd.Dir = constants.TmpDir
	cmd.Run()

	cmd = exec.Command("makepkg", "-si")
	cmd.Dir = constants.TmpDir
	cmd.Run()
}