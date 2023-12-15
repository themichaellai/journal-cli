package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a journal entry",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createFileWithCurrentDate(); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func createFileWithCurrentDate() error {
	fileName := time.Now().Format("2006-01-02")

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("file %s already exists", fileName)
		}
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	fmt.Printf("Created %s\n", fileName)
	return nil
}
