/*
Copyright © 2025 Wael Mahrous

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/waelmahrous/wormhole/internal"
)

var silent bool
var status bool
var Destination string
var FilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wormhole",
	Short: "Easily transport files between shells.",
	Long:  `Easily transport files between shells.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if silent {
			log.SetOutput(io.Discard)
		}

		FilePath = filepath.Join(Destination, ".wormhole.json")

		if err := os.MkdirAll(Destination, 0o755); err != nil {
			log.Printf("Could not create state directory %q: %v\n", Destination, err)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		if !status {
			_ = cmd.Help()
			return
		}

		dest, err := internal.GetDestination(FilePath)
		if err != nil {
			log.Printf("No open wormhole: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(dest)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Println("Could not establish user home directory")
		os.Exit(1)
	}

	rootCmd.Flags().BoolVarP(&status, "status", "t", false, "Show open wormhole")

	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Disable output")
	rootCmd.PersistentFlags().StringVarP(
		&Destination,
		"destination",
		"d",
		userHome,
		"Directory for wormhole state file",
	)
}
