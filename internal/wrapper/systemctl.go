package wrapper

import (
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
)

type Now bool

const EnableNow Now = true
const EnableOnly Now = false

func SystemctlEnable(service string, now Now) {
	out, err := exec.Command("systemctl", "is-active", service).Output()
	if err != nil {
		if err.Error() != "exit status 3" {
			pterm.Fatal.Println("Failed to check:", err)
		}
	}

	if strings.TrimSpace(string(out)) == "inactive" {
		pterm.Debug.Println("Enabling", service, "service...")
		var err error

		if now {
			err = Run("sudo", "systemctl", "--now", "enable", service)
		} else {
			err = Run("sudo", "systemctl", "enable", service)
		}

		if err != nil {
			pterm.Fatal.Println("Failed to start docker service:", err)
		}
	}
}
