/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
	"github.com/spf13/cobra"

	"github.com/waelmahrous/wormhole/internal"
)

var copyMode bool

func Transfer(src []string, dst string) ([]string, error) {
	if len(src) < 1 {
		return nil, errors.New("no files to send")
	}

	output := []string{}

	for _, v := range src {
		filePath := filepath.Join(filepath.Join(dst, filepath.Base(v)))
		if err := copy.Copy(v, filePath); err != nil {
			return output, err
		}

		output = append(output, filePath)

		if copyMode {
			continue
		}

		if err := os.Remove(v); err != nil {
			return output, err
		}
	}

	return output, nil
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [files...]",
	Short: "Send files into the active wormhole",
	Args:  cobra.MinimumNArgs(1), // require at least one file

	Run: func(cmd *cobra.Command, args []string) {
		target, err := internal.GetDestination(StateDir)
		if err != nil {
			internal.Fatalf("No open wormhole: %v\n", err)
		}

		log.Println("sending", len(args), "file(s) to", target)

		if _, err := Transfer(args, target); err != nil {
			internal.Fatalf("Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolVarP(&copyMode, "copy", "c", false, "Copy mode (do not move files)")
}
