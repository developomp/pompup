package bootstrap

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/etc/pacman.conf
var pacmanConf string

func setupPacman() {
	wrapper.SudoWriteFile("/etc/pacman.conf", pacmanConf)
}
