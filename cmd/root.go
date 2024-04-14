package cmd

import (
	"os"
	"slices"
	"strings"

	"github.com/developomp/pompup/internal/bootstrap"
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
	Long: `pompup is a personal one-click Arch Linux desktop setup utility tailor made for myself.

GitHub: https://github.com/developomp/pompup`,

	Run: func(cmd *cobra.Command, args []string) {
		cleanup()
		bootstrap.Bootstrap()

		var reminders []string = []string{
			"Reboot the system and run pompup again to complete installation",
		}

		// sort installers by name
		slices.SortStableFunc(installers.Installers, func(a, b *installers.Installer) int {
			if a.Priority != b.Priority {
				return a.Priority - b.Priority
			}

			return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
		})

		// run installers
		total := len(installers.Installers)
		for i, installer := range installers.Installers {
			pterm.Info.Printfln(
				"[%v / %v] %v - %v",
				i+1,
				total,
				pterm.FgWhite.Sprint(installer.Name),
				pterm.FgGray.Sprint(installer.Desc),
			)
			reminders = append(reminders, installer.Reminders...)
			installer.Setup()
		}

		// show reminders
		pterm.Info.Println("")
		pterm.Info.Println("Reminders:")
		for _, reminder := range reminders {
			pterm.Info.Println(reminder)
		}

		cleanup()
		pterm.Debug.Println("Done!")
	},
}

func cleanup() {
	// remove temporary directory if it exists
	if _, err := os.Stat(wrapper.TmpDir); err == nil {
		err := wrapper.Run("rm", "-rf", wrapper.TmpDir)
		if err != nil {
			pterm.Fatal.Printfln("Failed to clean '%s': %s", wrapper.TmpDir, err)
		}
	}
}
