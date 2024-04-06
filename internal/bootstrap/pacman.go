package bootstrap

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/etc/pacman.conf
var pacmanConf string

func setupPacman() {
	const configPath = "/etc/pacman.conf"
	if !wrapper.IsFileUpdated(configPath, pacmanConf) {
		wrapper.SudoWriteFile(configPath, pacmanConf)
	}
}
