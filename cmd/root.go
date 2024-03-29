package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile   string
	hostsfile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghosts",
	Short: "A host editing tool",
	Long:  `A host editing tool for a finer generation`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&hostsfile, "file", "f", "/etc/hosts", "Hosts file to read from")
}
