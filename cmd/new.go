package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	editor string // Define a flag variable to store the editor choice
)

func CreateAndAddTemplate(src string, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file: ", err)
		return
	}
	defer srcFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying file contents:", err)
		return
	}

	// Check if the editor flag is set and open the file with the specified editor
	if editor != "" {
		openWithEditor(dest)
	}
}

func ExecuteCommand(cmd *cobra.Command, args []string) {
	filename := args[0]
	extSlice := strings.Split(filename, ".")
	if len(extSlice) != 2 {
		fmt.Println("Invalid filename!!")
		return
	}
	ext := extSlice[1]
	switch ext {
	case "go":
		CreateAndAddTemplate(fmt.Sprintf("/home/karthik/GolandProjects/tem/templates/%v.txt", ext), filename)
	default:
		fmt.Println("Sorry, the specified format is not yet supported!!")
	}
}

// Function to open a file with the specified editor
func openWithEditor(file string) {
	cmd := exec.Command(editor, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening file with editor:", err)
	}
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create and add templated to program files",
	Args:  cobra.ExactArgs(1),
	Run:   ExecuteCommand,
}

func init() {
	// Add a flag to NewCmd to specify the editor
	NewCmd.Flags().StringVarP(&editor, "editor", "e", "", "Specify the editor to open the file with (e.g., vscode)")
}
