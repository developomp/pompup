package bootstrap

import (
	"os"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

func Bootstrap() {
	pterm.Debug.Println("Bootstrapping...")

	// check for irrecoverable issues
	if err := wrapper.StartupCheck(); err != nil {
		pterm.Fatal.Println("Failed to start:", err)
	}

	// create temporary directory
	if err := os.MkdirAll(wrapper.TmpDir, wrapper.DefaultDirPerm); err != nil {
		pterm.Fatal.Println(err)
	}

	// install stuff
	setupPacman()
	setupParu()
	setupBash()
	setupZSH()
	wrapper.PacmanOnce("iputils")
	wrapper.PacmanOnce("wget")
	wrapper.PacmanOnce("trash-cli")
	wrapper.PacmanOnce("flatpak")

	pterm.Debug.Println("Bootstrapped!")
}
