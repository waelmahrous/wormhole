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

func transfer(src, dst string) error {
	if copyMode {
		return copy.Copy(src, dst)
	}
	return os.Rename(src, dst)
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		target, err := internal.GetDestination(FilePath)
		if err != nil {
			internal.Fatalf("No open wormhole: %v\n", err)
		}

		log.Println("sending", len(args), "file(s) to", target)

		for _, name := range args {
			src := name
			dst := filepath.Join(target, name)

			if err := transfer(src, dst); err != nil {
				internal.Fatalf("Failed to transfer %q to %q: %v\n", src, dst, err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode (do not move files)")
}
