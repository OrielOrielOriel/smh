/*
Copyright © 2021 Oriel <Orianafarrugia3@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/orielorieloriel/smh/libsmh"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smh",
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().DurationP("delay", "d", 0, "Time waited between requests, global across all jobs. Threads wait for each other. (e.g. 1500ms")
	rootCmd.PersistentFlags().Bool("no-error", false, "Don't display errors.")
	rootCmd.PersistentFlags().Bool("no-progress", false, "Don't display progress.")
	rootCmd.PersistentFlags().Bool("no-status", false, "Don't display rolling request status.")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file to write results to (defaults to stdout), affected by verbosity settings.")
	rootCmd.PersistentFlags().Bool("quiet", false, "Don't print the banner and other noise.")
	rootCmd.PersistentFlags().IntP("threads", "t", 10, "Number of concurrent threads.")
	rootCmd.PersistentFlags().CountP("verbose", "v", 0, "Verbosity level. Goes up to _")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".smh" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".smh")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
