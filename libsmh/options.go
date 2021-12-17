package libsmh

import "time"

// Options holds all the options that could be passed to libsmh
type Options struct {
	Delays         time.Duration
	NoError        bool
	NoProgress     bool
	NoStatus       bool
	OutputFilename string
	Quiet          bool
	Threads        int
	Verbose        bool
	Wordlists      map[string]string
}

// JobOptions holds all the options that can be passed to a libsmh.Job
type JobOptions struct {
	Delays   map[string]time.Duration
	Lockout  int
	Protocol string
}

// NewOptions returns a new initialized Options object
func NewOptions() *Options {
	return &Options{}
}

// NewOptions returns a new initialized Options object
func NewJobOptions() *JobOptions {
	return &JobOptions{}
}
