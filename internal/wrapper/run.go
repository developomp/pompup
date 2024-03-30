package wrapper

import (
	"os"
	"os/exec"
)

func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	// show error messages
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// BashRun runs a command as if it was typed out in a bash shell.
// This allows the usage of piping among other stuff.
func BashRun(command string) error {
	return Run("bash", "-c", command)
}
