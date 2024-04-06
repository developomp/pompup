package wrapper

import (
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
)

func SystemctlEnable(service string) {
	out, err := exec.Command("systemctl", "is-active", service).Output()
	if err != nil {
		if err.Error() != "exit status 3" {
			pterm.Fatal.Println("Failed to check:", err)
		}
	}

	if strings.TrimSpace(string(out)) == "inactive" {
		pterm.Debug.Println("Enabling", service, "service...")
		err = Run("sudo", "systemctl", "--now", "enable", service)
		if err != nil {
			pterm.Fatal.Println("Failed to start docker service:", err)
		}
	}
}
