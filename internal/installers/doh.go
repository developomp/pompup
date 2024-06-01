package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/etc/dns-over-https/doh-client.conf
var _dohConf string

// https://wiki.archlinux.org/title/DNS-over-HTTPS
func init() {
	register(&Installer{
		Name: "DoH",
		Desc: "DNS over HTTPS",
		Setup: func() {
			wrapper.ParuOnce("dns-over-https")
			wrapper.SystemctlEnable("doh-client")

			wrapper.SudoWriteFile("/etc/dns-over-https/doh-client.conf", _dohConf)
		},
		Reminders: []string{
			"Perform https://wiki.archlinux.org/title/DNS-over-HTTPS#Troubleshooting",
		},
	})
}
