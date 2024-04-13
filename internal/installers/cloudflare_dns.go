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
		Setup: func() {
			const filePath = "/etc/resolv.conf"

			if wrapper.IsFileUpdated(filePath, _resolvConf) {
				return
			}

			// remove immutable flag from /etc/resolv.conf
			err := wrapper.Run("sudo", "chattr", "-i", filePath)
			if err != nil {
				pterm.Fatal.Println("Failed to remove immutable flag from /etc/resolv.conf:", err)
			}

			// write to /etc/resolv.conf
			wrapper.SudoWriteFile(filePath, _resolvConf)

			// restore immutable flag
			err = wrapper.Run("sudo", "chattr", "+i", filePath)
			if err != nil {
				pterm.Fatal.Println("Failed to restore immutable flag of /etc/resolv.conf:", err)
			}
		},
	})
}
