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
	Spare      int  // spare the last amount of GB from filling up
	DataPoints int  // number of data points for speed measurement
	Wait       int  // wait time after write before read, in seconds
	KbSpeed    bool // use KB/s to measure speed
	Debug      int  // debugging level
	Help       bool // show this usage help
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
	flag.IntVar(&Opts.DataPoints, "p", 100,
		"number of data points for speed measurement")
	flag.IntVar(&Opts.Wait, "w", 121,
		"wait time after write before read, in seconds")
	flag.BoolVar(&Opts.KbSpeed, "k", false,
		"use KB/s to measure speed")
	flag.IntVar(&Opts.Debug, "d", 0,
		"debugging level")
	flag.BoolVar(&Opts.Help, "h", false,
		"show this usage help")
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
	if Opts.DataPoints == 0 &&
		len(os.Getenv("DISKCHECK_P")) != 0 {
		if i, err := strconv.Atoi(os.Getenv("DISKCHECK_P")); err == nil {
			Opts.DataPoints = i
		}
	}
	if Opts.Wait == 0 &&
		len(os.Getenv("DISKCHECK_W")) != 0 {
		if i, err := strconv.Atoi(os.Getenv("DISKCHECK_W")); err == nil {
			Opts.Wait = i
		}
	}
	if _, exists = os.LookupEnv("DISKCHECK_K"); Opts.KbSpeed || exists {
		Opts.KbSpeed = true
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

const usageSummary = "  -sp\tspare the last amount of GB from filling up (DISKCHECK_SP)\n  -p\tnumber of data points for speed measurement (DISKCHECK_P)\n  -w\twait time after write before read, in seconds (DISKCHECK_W)\n  -k\tuse KB/s to measure speed (DISKCHECK_K)\n  -d\tdebugging level (DISKCHECK_D)\n  -h\tshow this usage help (DISKCHECK_H)\n\nDetails:\n\n"

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

The '-sp','-p' flags can be overridden by environment variables
'DISKCHECK_SP','DISKCHECK_P', etc

Usage Examples (say disk avail is 267GB):
  diskcheck -sp 203 /tmp # 64GB write
  diskcheck -sp 139 -p 200 /tmp # 128GB write
  ... # 256GB write with -p 400

  or,

  DISKCHECK_SP=139 DISKCHECK_P=200 diskcheck /tmp
`)
	os.Exit(0)
}
