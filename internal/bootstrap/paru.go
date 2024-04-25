package bootstrap

import (
	_ "embed"

	"os"
	"os/exec"
	"path/filepath"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/etc/paru.conf
var paruConf string

func setupParu() {
	if !wrapper.IsArchPkgInstalled("paru-bin") {
		installParu()
	}

	wrapper.SudoWriteFile("/etc/paru.conf", paruConf)
}

func installParu() {
	wrapper.PacmanOnce("git")
	wrapper.PacmanOnce("base-devel")

	var cmd *exec.Cmd

	cmd = exec.Command("git", "clone", "https://aur.archlinux.org/paru-bin.git")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = wrapper.TmpDir
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Println("Failed to clone https://aur.archlinux.org/paru-bin.git:", err)
	}

	cmd = exec.Command("makepkg", "-si", "--noconfirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Join(wrapper.TmpDir, "paru-bin")
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Println("Failed to install paru:", err)
	}
}
