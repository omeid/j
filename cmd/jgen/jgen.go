package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	output    = flag.String("output", "", "output file name; default srcdir/<type>_j.go")
)

// Usage is a replacement usage function for the flags package.
func Usage() {

	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tjgen [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tjgen [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttp://github.com/omeid/j\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	var err error
	if err != nil {
		return
	}
}
