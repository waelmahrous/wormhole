/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a wormhole in the current directory",

	Run: func(cmd *cobra.Command, args []string) {
		target, err := os.Getwd()

		if err != nil {
			log.Println("Could not open wormhole in target directory")
			os.Exit(1)
		}

		err = os.WriteFile(FilePath, []byte(target), 0x644)
		if err != nil {
			log.Println("Could not set target directory")
			os.Exit(1)
		}

		log.Println("Wormhole open at", target)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
