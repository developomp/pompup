package cmd

import (
	"os"

	"github.com/developomp/pompup/internal/constants"
	"github.com/developomp/pompup/internal/installers"
	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

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
	Long: `pompup is a personal Arch Linux desktop setup utility tailor-made for developomp.

GitHub: https://github.com/developomp/pompup`,

	Run: func(cmd *cobra.Command, args []string) {
		initialize()
		defer cleanup()

		var reminders []string

		// run setup functions
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

func initialize() {
	pterm.Debug.Println("Initializing...")
	defer pterm.Debug.Println("Initialized!")

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

	// install dependencies
	wrapper.Deps()

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
