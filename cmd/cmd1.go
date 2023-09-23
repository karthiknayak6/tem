package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmd1Cmd = &cobra.Command{
    Use:   "cmd1",
    Short: "A brief description of cmd1",
    Args:  cobra.ExactArgs(1), // Accepts exactly one argument
    Run: func(cmd *cobra.Command, args []string) {
        arg := args[0]
        fmt.Printf("Executing cmd1 with argument: %s\n", arg)
    },
}

func init() {
    // Add flags for cmd1 here, if needed
}

// Export the cmd1Cmd variable to make it available to main.go
var Cmd1 = cmd1Cmd
