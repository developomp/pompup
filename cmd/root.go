package cmd

import (
	"log"
	"os"

	"github.com/developomp/pompup/internal/check"
	"github.com/developomp/pompup/internal/constants"
	"github.com/developomp/pompup/internal/install"
	"github.com/developomp/pompup/internal/ui"
	"github.com/developomp/pompup/internal/workflows"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Failed to run cmd.Execute()")
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

		var postSetups []func()

		// run setup functions
		for i, selected := range ui.Select() {
			if !selected {
				continue
			}

			workflow := workflows.Workflows[i]
			postSetups = append(postSetups, workflow.PostSetup)
			workflow.Setup()
		}

		// run post-setup functions
		for _, postSetup := range postSetups {
			if postSetup == nil {
				continue
			}
			postSetup()
		}

	},
}

func initialize() {
	log.Println("Initializing...")
	defer log.Println("Initialized!")

	// perform startup checks
	if err := check.StartupCheck(); err != nil {
		log.Fatalln("Failed to start:", err)
	}

	// install dependencies
	install.Deps()

	// create temporary directory
	if err := os.MkdirAll(constants.TmpDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
}

func cleanup() {
	log.Println("Wrapping up...")
	defer log.Println("Done!")

	// remove temporary directory
	if err := os.RemoveAll(constants.TmpDir); err != nil {
		log.Fatalln(err)
	}
}
