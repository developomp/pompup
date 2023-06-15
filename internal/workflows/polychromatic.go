package workflows

import (
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/install"
)

func init() {
	register(&Workflow{
		Name: "Dconf Editor",
		Desc: "GSettings editor for GNOME",
		Tags: []Tag{Configurator},
		Setup: func() {
			install.Paru("openrazer-meta")
			exec.Command("sudo", "gpasswd", "-a", os.Getenv("USER"), "plugdev").Run()
			install.Paru("polychromatic")
		},
	})
}
