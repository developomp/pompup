package workflows

import (
	_ "embed"

	"github.com/developomp/pompup/internal/helper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/resolv.conf
var _resolv_conf string

func init() {
	register(&Workflow{
		Name: "Cloudflare DNS",
		Desc: "1.1.1.1 DNS",
		Tags: []Tag{System},

		Setup: func() {
			// remove immutable flag from /etc/resolv.conf
			err := helper.Run("sudo", "chatty", "-i", "/etc/resolv.conf")
			if err != nil {
				pterm.Fatal.Println("Failed to remove immutable flag from /etc/resolv.conf:", err)
			}

			// write to /etc/resolv.conf
			err = helper.SudoWriteFile("/etc/resolv.conf", _resolv_conf)
			if err != nil {
				pterm.Fatal.Println("Failed to write /etc/resolv.conf:", err)
			}

			// restore immutable flag
			err = helper.Run("sudo", "chattr", "+i", "/etc/resolv.conf")
			if err != nil {
				pterm.Fatal.Println("Failed to restore immutable flag of /etc/resolv.conf:", err)
			}
		},
	})
}
