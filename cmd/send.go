/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("sending", len(args), "file(s)")

		target, _ := os.ReadFile(FilePath)

		for _, fileName := range args {
			os.Rename(fileName, filepath.Join(string(target), fileName))
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
