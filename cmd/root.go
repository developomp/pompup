package cmd

import (
	_ "embed"
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/constants"
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
		bootstrap()
		defer cleanup()

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
	},
}

func bootstrap() {
	pterm.Debug.Println("Bootstrapping...")
	defer pterm.Debug.Println("Bootstrapped!")

	// perform startup checks
	if err := wrapper.StartupCheck(); err != nil {
		pterm.Fatal.Println("Failed to start:", err)
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

	// create temporary directory
	if err := os.MkdirAll(constants.TmpDir, wrapper.DefaultPerm); err != nil {
		pterm.Fatal.Println(err)
	}
}

func cleanup() {
	pterm.Debug.Println("Wrapping up...")
	defer pterm.Debug.Println("Done!")

	// remove temporary directory
	if err := os.RemoveAll(constants.TmpDir); err != nil {
		pterm.Fatal.Printfln("Failed to clean '%s': %s", constants.TmpDir, err)
	}
}

func installParu() {
	wrapper.Pacman("git")
	wrapper.Pacman("base-devel")

	var cmd *exec.Cmd

	cmd = exec.Command("git", "clone", "https://aur.archlinux.org/paru.git")
	cmd.Stderr = os.Stderr
	cmd.Dir = constants.TmpDir
	cmd.Run()

	cmd = exec.Command("makepkg", "-si")
	cmd.Stderr = os.Stderr
	cmd.Dir = constants.TmpDir
	cmd.Run()

	wrapper.SudoWriteFile("/etc/pacman.conf", pacmanConf)
	wrapper.SudoWriteFile("/etc/paru.conf", paruConf)
}
