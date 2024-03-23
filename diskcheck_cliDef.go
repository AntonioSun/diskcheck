// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "diskcheck" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Spare int  // spare the last amount of GB from filling up
	Debug int  // debugging level
	Help  bool // show usage help
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func initVars() {

	// set default values for command line parameters
	flag.IntVar(&Opts.Spare, "sp", 2,
		"spare the last amount of GB from filling up")
	flag.IntVar(&Opts.Debug, "d", 0,
		"debugging level")
	flag.BoolVar(&Opts.Help, "h", false,
		"show usage help")
}

func initVals() {
	exists := false
	// Now override those default values from environment variables
	if Opts.Spare == 0 &&
		len(os.Getenv("DISKCHECK_SP")) != 0 {
		if i, err := strconv.Atoi(os.Getenv("DISKCHECK_SP")); err == nil {
			Opts.Spare = i
		}
	}
	if Opts.Debug == 0 &&
		len(os.Getenv("DISKCHECK_D")) != 0 {
		if i, err := strconv.Atoi(os.Getenv("DISKCHECK_D")); err == nil {
			Opts.Debug = i
		}
	}
	if _, exists = os.LookupEnv("DISKCHECK_H"); Opts.Help || exists {
		Opts.Help = true
	}

}

const usageSummary = "  -sp\tspare the last amount of GB from filling up (DISKCHECK_SP)\n  -d\tdebugging level (DISKCHECK_D)\n  -h\tshow usage help (DISKCHECK_H)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags..] writable_path\n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, usageSummary)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		`
The program will fill up the remaining of disk space given by
the 'writable_path', and leave the last 'spare' amount of GB
free for normal operation.

The '-sp' flag can be overridden by environment variable(s)
'DISKCHECK_SP'
`)
	os.Exit(0)
}
