package bootstrap

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.bashrc
var bashConfig []byte

func setupBash() {
	wrapper.WriteFile(wrapper.InHome(".bashrc"), bashConfig)
}
