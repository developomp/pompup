package cmd

import (
	"os"

	"github.com/developomp/pompup/internal/check"
	"github.com/developomp/pompup/internal/constants"
	"github.com/developomp/pompup/internal/helper"
	"github.com/developomp/pompup/internal/install"
	"github.com/developomp/pompup/internal/ui"
	"github.com/developomp/pompup/internal/workflows"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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
		for i, selected := range ui.Select() {
			if !selected {
				continue
			}

			workflow := workflows.Workflows[i]
			reminders = append(reminders, workflow.Reminders...)
			workflow.Setup()
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
	if err := check.StartupCheck(); err != nil {

		if err == check.ErrStartupNoWget {
			install.Pacman("wget")
		} else {
			pterm.Fatal.Println("Failed to start:", err)
		}
	}

	// install dependencies
	install.Deps()

	// create temporary directory
	if err := os.MkdirAll(constants.TmpDir, helper.DefaultPerm); err != nil {
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
