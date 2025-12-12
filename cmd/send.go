/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"

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
		target, err := internal.GetDestination(StateDir)
		if err != nil {
			log.Fatalf("No open wormhole: %v\n", err)
		}

		log.Println("sending", len(args), "file(s) to", target)

		record := internal.TransferRecord{
			Source:      args,
			Copy:        copyMode,
			StateDir:    StateDir,
		}

		if _, err := internal.Transfer(record); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode (do not move files)")
}
