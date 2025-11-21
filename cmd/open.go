/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a wormhole in the current directory",

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Opening wormhole...")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
