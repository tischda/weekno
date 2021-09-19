package main

import (
	"fmt"
	"time"
)

// set via -ldflags
var version string

func main() {
	_, week := time.Now().ISOWeek()
	fmt.Printf("%d", week)
}
