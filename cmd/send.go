/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
	"github.com/spf13/cobra"

	"github.com/waelmahrous/wormhole/internal"
)

var copyMode bool

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		target, err := internal.GetDestination(FilePath)
		if err != nil {
			log.Printf("No open wormhole: %v\n", err)
			os.Exit(1)
		}

		log.Println("sending", len(args), "file(s) to", target)

		for _, fileName := range args {
			dst := filepath.Join(target, fileName)

			if copyMode {
				if err := copy.Copy(fileName, dst); err != nil {
					log.Printf("Failed to copy %q to %q: %v\n", fileName, dst, err)
					os.Exit(1)
				}
			} else {
				if err := os.Rename(fileName, dst); err != nil {
					log.Printf("Failed to move %q to %q: %v\n", fileName, dst, err)
					os.Exit(1)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode (do not move files)")
}
