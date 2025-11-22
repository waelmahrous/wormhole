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
	"path/filepath"

	"github.com/spf13/cobra"
)

var verbose bool
var Destination string
var FilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wormhole",
	Short: "Easily transport files between shells.",
	Long:  `Easily transport files between shells.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose == false {
			log.SetOutput(io.Discard)
		}

		FilePath = filepath.Join(Destination, ".wormhole.state")
	},

	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(FilePath)
		if err == nil {
			log.Println("State file already exists at ", FilePath)
		} else if os.IsNotExist(err) {
			_, err := os.Create(FilePath)

			if err != nil {
				log.Println("Could not create state file")
				os.Exit(1)
			}

			log.Println("Created state file at ", Destination)
		} else {
			log.Println("Error: ", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var userHome, err = os.UserHomeDir()

	if err != nil {
		log.Println("Could not establish user home directory")
		os.Exit(1)
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&Destination, "destination", "d", userHome, "Set custom state file directory")
}
