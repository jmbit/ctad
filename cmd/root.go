/*
Copyright © 2024 Johannes Bülow <johannes.buelow@jmbit.de>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jmbit/ctad/playback"
	"github.com/spf13/cobra"
)

var cfgFile string
var soundDir string
var maxDelay = 500
var minDelay = 30

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ctad",
	Short: "Randomly plays cat sounds",
	Long: `Program that lets your computer meow (or play any other noise, really), in a random interval
  the delay set is in Seconds. 
  Sound files have to be mp3 and all in a single directory, 3 cat sounds are already embedded
  Example:
    ctad --max-delay 60 --min-delay 10 -s ./sounds/
  plays mp3 files from ./sounds every 10-30s`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting ctad...")
    for {
      // Get random delay. If it's lower than the configured minimum, use that instead
      delay := rand.Intn(maxDelay)
      if delay < minDelay {
        log.Printf("Sleeping for %d seconds", minDelay)
        time.Sleep(time.Duration(minDelay)*time.Second)
      } else {
        log.Printf("Sleeping for %d seconds", delay)
        time.Sleep(time.Duration(delay)*time.Second)
      }
      playback.PlayRandomSound(soundDir)
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
	log.SetOutput(os.Stderr)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&soundDir, "sounds", "s", "", "directory to use for sound files")
  rootCmd.PersistentFlags().IntVarP(&maxDelay, "max-delay", "m", 500, "Maximum delay between noises")
  rootCmd.PersistentFlags().IntVarP(&minDelay, "min-delay", "i", 30, "Maximum delay between noises")
}
