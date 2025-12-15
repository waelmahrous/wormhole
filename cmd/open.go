/*
Copyright © 2025 Wael Mahrous
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var destination string

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a wormhole in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			target string
			err    error
		)

		if destination != "" {
			target = destination
		} else {
			if target, err = os.Getwd(); err != nil {
				log.Fatalf("Could not resolve working directory %v\n", err)
			}
		}

		if wormhole, err := Wormhole.SetDestination(target); err != nil {
			log.Fatalf("Could not open wormhole in %q: %v\n", target, err)
		} else {
			log.Printf("Wormhole open at %s", wormhole.Destination)
		}
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
