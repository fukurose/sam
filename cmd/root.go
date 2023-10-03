package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sam",
	Short: "file porter within the local network",
	Long:  "Sam is a file porter within the local network.",
}

func init() {
	rootCmd.PersistentFlags().StringP("address", "a", ":50000", "Server Address")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
