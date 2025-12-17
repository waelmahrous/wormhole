/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/waelmahrous/wormhole/internal"
)

var (
	copyMode bool
	force    bool
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		target, err := Wormhole.GetDestination()
		if err != nil {
			log.Fatalf("No open wormhole: %v\n", err)
		}

		log.Println("sending", len(args), "file(s) to", target)

		record := internal.TransferRecord{
			Source:     args,
			Copy:       copyMode,
			WormholeID: Wormhole.ID,
			Force:      force,
		}

		if SafeMode {
			safeZonePath := filepath.Join(Wormhole.StateDir, internal.DefaultSafeZone)
			os.Mkdir(safeZonePath, 0755)

			safeZone := internal.Wormhole{
				ID:          "safezone",
				Destination: safeZonePath,
				StateDir:    StateDir,
			}

			if err := safeZone.InitWormholeStore(); err != nil {
				log.Fatal(err)
			}

			backupRecord := record
			backupRecord.Copy = true
			backupRecord.WormholeID = safeZone.ID

			if _, err := safeZone.Transfer(backupRecord); err != nil {
				log.Fatalf("Could not copy files to safezone, %v", err)
			}

			log.Printf("copied %d file(s) to safezone in %s", len(args), safeZone.Destination)
		}

		if _, err := Wormhole.Transfer(record); err != nil {
			log.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode (do not move files)")
	sendCmd.Flags().BoolVarP(&force, "force", "f", false, "Force send, if file already exists")
}
