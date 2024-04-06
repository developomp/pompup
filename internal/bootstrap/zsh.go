package bootstrap

import (
	_ "embed"
	"os"
	"os/exec"
	"strings"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.zshrc
var zshConfig []byte

func setupZSH() {
	wrapper.ParuOnce("zsh")
	installOMZ()
	installP10K()

	err := wrapper.WriteFile(wrapper.InHome(".zshrc"), zshConfig)
	if err != nil {
		pterm.Fatal.Println("Failed to restore zsh config file:", err)
	}

	// $SHELL could be /usr/bin/zsh
	if !strings.HasSuffix(os.Getenv("SHELL"), "/bin/zsh") {
		// set the default shell to zsh
		if err := exec.Command("sudo", "chsh", "-s", "/bin/zsh").Run(); err != nil {
			pterm.Fatal.Println("Failed to set default shell to zsh")
		}
	}
}

func installOMZ() {
	// install OMZ if it is not installed already
	if !wrapper.PathExists(wrapper.InHome(".oh-my-zsh")) {
		pterm.Debug.Println("Installing https://github.com/ohmyzsh/ohmyzsh")

		cmd := exec.Command("sh", "-c", "sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/wrapper.sh)\"")
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			pterm.Fatal.Println("Failed to install Oh My Zsh:", err)
		}
	}

	// install syntax highlighter plugin only if it is not installed already
	pluginPath := wrapper.InHome(".oh-my-zsh/custom/plugins/zsh-syntax-highlighting")
	if !wrapper.PathExists(pluginPath) {
		pterm.Debug.Println("Installing zsh syntax highlighter")

		if err := exec.Command("git", "clone", "--depth=1", "https://github.com/zsh-users/zsh-syntax-highlighting.git", pluginPath).Run(); err != nil {
			pterm.Fatal.Println("Failed to install https://github.com/zsh-users/zsh-syntax-highlighting:", err)
		}
	}
}

func installP10K() {
	p10kPath := wrapper.InHome(".oh-my-zsh/custom/themes/powerlevel10k")

	if wrapper.PathExists(p10kPath) {
		return
	}

	pterm.Debug.Println("Installing p10k theme")
	if err := exec.Command("git", "clone", "--depth=1", "https://github.com/romkatv/powerlevel10k.git", p10kPath).Run(); err != nil {
		pterm.Fatal.Printfln("Failed to clone https://github.com/romkatv/powerlevel10k.git: %s", err)
	}
}
