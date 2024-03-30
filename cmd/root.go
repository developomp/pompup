package cmd

import (
	_ "embed"
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/installers"
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

//go:embed assets/etc/pacman.conf
var pacmanConf string

//go:embed assets/etc/paru.conf
var paruConf string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	pterm.EnableDebugMessages() // allows pterm..Debug.*() to be used

	if err := rootCmd.Execute(); err != nil {
		pterm.Fatal.Println("Failed to run cmd.Execute()")
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "pompup",
	Long: `pompup is a personal one-click Arch Linux desktop setup utility tailor made for myself.

GitHub: https://github.com/developomp/pompup`,

	Run: func(cmd *cobra.Command, args []string) {
		cleanup()
		bootstrap()

		var reminders []string

		// run installers
		for _, installer := range installers.Installers {
			reminders = append(reminders, installer.Reminders...)
			installer.Setup()
		}

		// show reminders
		for _, reminder := range reminders {
			pterm.Info.Println(reminder)
		}

		pterm.Debug.Println("Wrapping up...")
		cleanup()
		pterm.Debug.Println("Done!")
	},
}

func bootstrap() {
	pterm.Debug.Println("Bootstrapping...")

	// check for irrecoverable issues
	if err := wrapper.StartupCheck(); err != nil {
		pterm.Fatal.Println("Failed to start:", err)
	}

	// create temporary directory
	if err := os.MkdirAll(wrapper.GetTmpDir(), wrapper.DefaultDirPerm); err != nil {
		pterm.Fatal.Println(err)
	}

	if !wrapper.IsBinInstalled("ping") {
		wrapper.Pacman("iputils")
	}

	if !wrapper.IsBinInstalled("wget") {
		wrapper.Pacman("wget")
	}

	if !wrapper.IsBinInstalled("trash") {
		wrapper.Pacman("trash-cli")
	}

	if !wrapper.IsBinInstalled("flatpak") {
		wrapper.Pacman("flatpak")
	}

	if !wrapper.IsBinInstalled("paru") {
		installParu()
	}

	pterm.Debug.Println("Bootstrapped!")
}

func cleanup() {
	// remove temporary directory if it exists
	if _, err := os.Stat(wrapper.GetTmpDir()); err == nil {
		err := wrapper.Run("rm", "-rf", wrapper.GetTmpDir())
		if err != nil {
			pterm.Fatal.Printfln("Failed to clean '%s': %s", wrapper.GetTmpDir(), err)
		}
	}
}

func installParu() {
	if !wrapper.IsBinInstalled("git") {
		wrapper.Pacman("git")
	}

	if !wrapper.IsArchPkgInstalled("pacman", "base-devel") {
		wrapper.Pacman("base-devel")
	}

	var cmd *exec.Cmd

	cmd = exec.Command("git", "clone", "https://aur.archlinux.org/paru-bin.git")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = wrapper.GetTmpDir()
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Println("Failed to clone https://aur.archlinux.org/paru-bin.git:", err)
	}

	cmd = exec.Command("makepkg", "-si", "--noconfirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = wrapper.GetTmpDir() + "/paru-bin"
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Println("Failed to install paru:", err)
	}

	wrapper.SudoWriteFile("/etc/pacman.conf", pacmanConf)
	wrapper.SudoWriteFile("/etc/paru.conf", paruConf)
}
