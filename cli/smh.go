package cli

import (
	"fmt"
	"os"
	"time"
	
	"github.com/orielorieloriel/smh/libsmh"
)

const ruler = "==============================================================="
const cliProgressUpdate = 500 * time.Millisecond

func banner() {
	fmt.Printf("snake-many-head v%s\n", libsmh.VERSION)
	fmt.Println("by Oriel (@OrielOrielOriel)")
}

// SMH is the main entry point for the CLI
func SMH(opts *libsmh.Options, job_opts []*libsmh.JobOptions) error {
	// Sanity checks
	if opts == nil {
		return fmt.Errorf("valid global options not provided")
	}

	if job_opts == nil {
		return fmt.Errorf("valid job options not provided")
	}

	return nil
}
