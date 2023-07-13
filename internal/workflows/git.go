package workflows

import (
	"os/exec"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "git",
		Desc: "git gud",
		Tags: []Tag{Dev, Cli},
		Setup: func() {
			install.Paru("git")

			err := exec.Command("git", "config", "--global", "user.email", "developomp@gmail.com").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to set git email:", err)
			}

			err = exec.Command("git", "config", "--global", "user.name", "developomp").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to set git username:", err)
			}

			err = exec.Command("git", "config", "--global", "credential.helper", "store").Run()
			if err != nil {
				pterm.Fatal.Println("Failed to set git credential settings:", err)
			}
		},
	})
}
