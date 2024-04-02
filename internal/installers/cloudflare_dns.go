package installers

import (
	_ "embed"
	"os"
	"strings"

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
			const filePath = "/etc/resolv.conf"

			b, err := os.ReadFile(filePath)
			if err != nil {
				pterm.Fatal.Printfln("Failed to read %s: %s", filePath, err)
			}
			if strings.TrimSpace(string(b)) == _resolvConf {
				return
			}

			// remove immutable flag from /etc/resolv.conf
			err = wrapper.Run("sudo", "chattr", "-i", filePath)
			if err != nil {
				pterm.Fatal.Println("Failed to remove immutable flag from /etc/resolv.conf:", err)
			}

			// write to /etc/resolv.conf
			err = wrapper.SudoWriteFile(filePath, _resolvConf)
			if err != nil {
				pterm.Fatal.Println("Failed to write /etc/resolv.conf:", err)
			}

			// restore immutable flag
			err = wrapper.Run("sudo", "chattr", "+i", filePath)
			if err != nil {
				pterm.Fatal.Println("Failed to restore immutable flag of /etc/resolv.conf:", err)
			}
		},
	})
}
