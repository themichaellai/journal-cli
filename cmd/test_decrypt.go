package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testDecryptCmd = &cobra.Command{
	Use:   "test-decrypt",
	Short: "Test decryption",
	Long:  `Long description of the test decrypt subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test Decrypt command invoked")
		// Your test decryption logic goes here
	},
}
