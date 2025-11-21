/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a wormhole in the current directory",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open called")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
