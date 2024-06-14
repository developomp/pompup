package bootstrap

import (
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"

	_ "embed"
)

//go:embed assets/etc/nix/nix.conf
var nixConf string

//go:embed assets/home/.config/nixpkgs/config.nix
var configNix []byte

func setupNix() {
	if !wrapper.IsBinInstalled("nix-env") {
		pterm.Fatal.Println("Install Nix by running 'sh <(curl -L https://nixos.org/nix/install) --daemon'")
	}

	wrapper.SudoWriteFile("/etc/nix/nix.conf", nixConf)
	wrapper.WriteFile(wrapper.InHome(".config/nixpkgs/config.nix"), configNix)
	wrapper.ParuOnce("nix-zsh-completions")
}
