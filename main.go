package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	var Source string
	var isOn bool
	var versionCmd = &cobra.Command{
		Use:   "boarding",
		Short: "Print the version number go",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sub cmd version - 0.1", Source, isOn)
		},
	}

	versionCmd.Flags().StringVarP(&Source, "source", "s", "../../identity/default", "Source directory to read from")
	versionCmd.Flags().BoolVarP(&isOn, "on", "o", true, "Offboarding Employee(s), Default is Onboarding (WIP)")
	// versionCmd.MarkFlagRequired("source")

	var rootCmd = &cobra.Command{Use: "root"}
	rootCmd.AddCommand(versionCmd)

	var Verbose bool
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.Execute()

}
