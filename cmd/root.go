package cmd

import (
	"log"

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
	},
}
