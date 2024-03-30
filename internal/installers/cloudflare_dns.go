package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/resolv.conf
var _resolvConf string

func init() {
	register(&Installer{
		Name: "Cloudflare DNS",
		Desc: "1.1.1.1 DNS",
		Tags: []Tag{System},

		Setup: func() {
			// remove immutable flag from /etc/resolv.conf
			err := wrapper.Run("sudo", "chattr", "-i", "/etc/resolv.conf")
			if err != nil {
				pterm.Fatal.Println("Failed to remove immutable flag from /etc/resolv.conf:", err)
			}

			// write to /etc/resolv.conf
			err = wrapper.SudoWriteFile("/etc/resolv.conf", _resolvConf)
			if err != nil {
				pterm.Fatal.Println("Failed to write /etc/resolv.conf:", err)
			}

			// restore immutable flag
			err = wrapper.Run("sudo", "chattr", "+i", "/etc/resolv.conf")
			if err != nil {
				pterm.Fatal.Println("Failed to restore immutable flag of /etc/resolv.conf:", err)
			}
		},
	})
}
