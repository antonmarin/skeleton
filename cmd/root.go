package cmd

import (
	"github.com/antonmarin/skeleton/cmd/apply"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "skeleton",
	Short: "Bones for your project",
	Long:  `Tool for developers to start quickly.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(apply.Cmd)
}
