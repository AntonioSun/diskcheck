package main

//go:generate sh diskcheck_cliGen.sh

import (
	"flag"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// define flags
	initVars()
	// popoulate flag variables from ENV
	initVals()
	// popoulate flag variables from cli
	flag.Parse()
	if Opts.Help {
		Usage()
	}

	// There is one mandatory non-flag argument
	if len(flag.Args()) < 1 {
		Usage()
	}
	checkDisk(flag.Args()[0])
}

//==========================================================================
// support functions

func debug(input string, threshold int) {
	if !(Opts.Debug >= threshold) {
		return
	}
	print("] ")
	print(input)
	print("\n")
}

func checkError(err error) {
	if err != nil {
		log.Printf("%s: Fatal error - %s", progname, err.Error())
		os.Exit(1)
	}
}
