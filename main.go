package main

import (
	"fmt"
)

const (
	name = `Program`

	usage = "\n%s does nothing.\n\n" +
		"Usage: %s [OPTIONS] <args>...\n\nOPTIONS:\n"
)

// set via -ldflags
var version string

func main() {
	parseFlags()
	fmt.Println(name, " is doing nothing.")
}
