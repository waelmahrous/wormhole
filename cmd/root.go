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
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/waelmahrous/wormhole/internal"
)

var (
	silent      bool
	status      bool
	showVersion bool
	StateDir    string

	version = "dev"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wormhole",
	Short: "Easily transport files between shells.",
	Long:  `Easily transport files between shells.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := internal.InitWormholeStore(StateDir); err != nil {
			log.Fatal(err)
		}

		if silent {
			log.SetOutput(io.Discard)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		if status {
			if dest, err := internal.GetDestination(StateDir); err != nil {
				log.Fatalf("Could not get destination: %v\n", err)
			} else {
				log.Printf("Wormhole open in: %s", dest)
				return
			}
		}

		if showVersion {
			log.Printf("wormhole version: %s", version)
			return
		}

		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}

func init() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not establish user home directory")
	}

	rootCmd.Flags().BoolVarP(&status, "status", "t", false, "Show open wormhole")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Get wormhole version")

	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Disable output")
	rootCmd.PersistentFlags().StringVarP(
		&StateDir,
		"state-dir",
		"",
		userHome,
		"Directory where the wormhole state file is stored",
	)
}
