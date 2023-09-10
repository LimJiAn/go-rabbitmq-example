package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-rabbitmq-exam",
	Short: "CLI tool for rabbitmq exam",
	Long:  "go rabbitmq exam is a CLI tool for rabbitmq exam.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go [command]\n\n")
		fmt.Fprintf(os.Stderr, "Available commands are:\n")
		fmt.Fprintf(os.Stderr, "  publish (send)\n")
		fmt.Fprintf(os.Stderr, "  consume (receive)\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
