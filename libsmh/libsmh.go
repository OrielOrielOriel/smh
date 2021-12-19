package libsmh

import ()

type SMH struct {
	Opts *Options
	Jobs []*JobOptions
	RequestsExpected int
	RequestsIssued int
}
