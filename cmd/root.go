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

type ControlFlow = bool

var Continue ControlFlow = true
var Stop ControlFlow = false

var list bool = false
var only string = ""

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	pterm.EnableDebugMessages() // allows pterm..Debug.*() to be used

	rootCmd.PersistentFlags().BoolVar(&list, "list", list, "list available installers")
	rootCmd.PersistentFlags().StringVar(&only, "only", only, "only run specific installer")

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
		if list {
			listInstallers()
			os.Exit(0)
		}

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

		if only == "bootstrap" {

		} else if only != "" {
			for _, installer := range installers.Installers {
				if installer.Name != only {
					continue
				}

				pterm.Debug.Printfln("Found %s installer. Running...", installer.Name)
				reminders = append(reminders, installer.Reminders...)
				installer.Setup()
			}
		} else {
			// run installers
			loopInstallers(func(installer *installers.Installer, _ int, _ int) bool {
				reminders = append(reminders, installer.Reminders...)
				installer.Setup()

				return true
			})
			pterm.Info.Println("")
		}

		// show reminders
		pterm.Info.Println("Reminders:")
		for _, reminder := range reminders {
			pterm.Info.Println(reminder)
		}

		cleanup()
		pterm.Debug.Println("Done!")
	},
}

func listInstallers() {
	pterm.Info.Println("Listing available installers:")
	pterm.Info.Println("")

	loopInstallers(func(_ *installers.Installer, _ int, _ int) ControlFlow {
		return Continue
	})
}

func loopInstallers(f func(*installers.Installer, int, int) ControlFlow) {
	total := len(installers.Installers)
	for i, installer := range installers.Installers {
		pterm.Info.Printfln(
			"[%v / %v] %v - %v",
			i+1,
			total,
			pterm.FgWhite.Sprint(installer.Name),
			pterm.FgGray.Sprint(installer.Desc),
		)

		if !f(installer, i, total) {
			break
		}
	}
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
