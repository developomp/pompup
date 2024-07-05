package installers

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/etc/environment
var _etcEnvironment string

func init() {
	register(&Installer{
		Name: "Environment",
		Desc: "/etc/environment",
		Setup: func() {
			wrapper.SudoWriteFile("/etc/environment", _etcEnvironment)
		},
	})
}
