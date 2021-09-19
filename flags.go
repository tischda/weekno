package main

import (
	"flag"
	"fmt"
	"os"
)

var showVersion bool

func parseFlags() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, name, name)
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	} else if showVersion || flag.Arg(0) == "version" {
		fmt.Printf("%s version %s\n", name, version)
		os.Exit(0)
	}
}
