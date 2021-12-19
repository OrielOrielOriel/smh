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
	"os/signal"
	"log"
	"context"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smh",
	SilenceUsage: true,
}

var mainContext context.Context

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var cancel context.CancelFunc
	mainContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan:
			// caught CTRL+C
			fmt.Println("\n[!] Keyboard interrupt detected, terminating.")
			cancel()
		case <-mainContext.Done():
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// Leaving this in results in the same error appearing twice
		// Once before and once after the help output. Not sure if
		// this is going to be needed to output other errors that
		// aren't automatically outputted.
		// fmt.Println(err)
		os.Exit(1)
	}
}

// This has to be called as part of the pre-run for sub commands. Including
// this in the init() function results in the built-in `help` command not
// working as intended. The required flags should only be marked as required
// on the global flags when one of the non-help commands is used.
func configureGlobalOptions() {
	if err := rootCmd.MarkPersistentFlagRequired("wordlist"); err != nil {
		log.Fatalf("error on marking flag as required: %v", err)
	}
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
	rootCmd.PersistentFlags().CountP("verbose", "v", "Verbosity level. Goes up to _")

}

func initThreads() int {
	threads, err := rootCmd.Flags().GetInt("threads")
	
	if err != nil {
		fmt.Errorf("invalid value for threads: %w", err)
		os.Exit(1)
	}

	if threads <= 0 {
		fmt.Errorf("threads must be more than 0")
		os.Exit(1)
	}
	
	return threads
}

func parseGlobalOptions() (*libsmh.Options, error) {
	globalopts := libsmh.NewOptions()
	
	globalopts.Threads = initThreads()

	return globalopts, nil
}
