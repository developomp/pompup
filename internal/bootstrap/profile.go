package bootstrap

import (
	_ "embed"

	"github.com/developomp/pompup/internal/wrapper"
)

//go:embed assets/home/.profile
var _profile []byte

func setupProfile() {
	wrapper.WriteFile(wrapper.InHome(".profile"), _profile)
}
