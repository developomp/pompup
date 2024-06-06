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
	needsReLogin := false
	setupPacman()
	setupParu()
	setupBash()
	needsReLogin = setupZSH() || needsReLogin
	wrapper.PacmanOnce("wget")
	wrapper.PacmanOnce("flatpak")

	if needsReLogin {
		pterm.Info.Println("Log out and lock back to finish bootstrapping")
		os.Exit(0)
	}

	pterm.Debug.Println("Bootstrapped!")
}
