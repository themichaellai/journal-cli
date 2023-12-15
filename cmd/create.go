package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a journal entry",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create command invoked")
	},
}
