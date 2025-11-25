/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"github.com/otiai10/copy"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var copyMode bool

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := os.ReadFile(FilePath)
		log.Println("sending", len(args), "file(s) to ", string(target))

		if copyMode == true {
			for _, fileName := range args {
				copy.Copy(fileName, filepath.Join(string(target), fileName))
			}
		} else {
			for _, fileName := range args {
				os.Rename(fileName, filepath.Join(string(target), fileName))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode")
}
