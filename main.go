package main

import (
	"fmt"
	"os"

	"github.com/karthiknayak6/tem/cmd" // Import the subcommands package

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mycliapp",
    Short: "A brief description of your CLI app",
    // Define global flags and settings here
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // Add subcommands here
    rootCmd.AddCommand(cmd.Cmd1)
   
    // Add more subcommands as needed
}
