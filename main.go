package main

import (
	"fmt"
	"os"

	"github.com/karthiknayak6/tem/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tem",
	Short: "Create and add templated to program files",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmd.NewCmd)

}
