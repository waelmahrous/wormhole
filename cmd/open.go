/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/waelmahrous/wormhole/internal"
)

var destination string

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a wormhole in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		var err error

		if destination != "" {
			target = destination
			_, err = os.ReadDir(target)
		} else {
			target, err = os.Getwd()
		}

		if err != nil {
			internal.Fatalf("Could not open wormhole in target directory %q: %v\n", target, err)
		}

		if err := internal.UpdateDestination("", target); err != nil {
			internal.Fatalf("Could not save wormhole state for %q: %v\n", target, err)
		}

		log.Println("Wormhole open at", target)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(
		&destination,
		"destination",
		"d",
		"",
		"Open wormhole in custom destination",
	)
}
