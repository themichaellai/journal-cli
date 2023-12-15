package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{Use: "journal-cli"}
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(encryptCmd)
    rootCmd.AddCommand(testDecryptCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

