package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
)

func main() {
	// Define flags
	suffixFlag := flag.Bool("s", false, "Append the argument as a suffix instead of a prefix")
	// Custom usage message to integrate with flag parsing later
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-s] [prefix_or_suffix_string]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generates a ULID, optionally prepending or appending the given string.\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  prefix_or_suffix_string   String to prepend (default) or append (with -s) to the ULID.\n")
		fmt.Fprintf(os.Stderr, "                            If omitted, only the ULID is generated.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Seed the random number generator
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	// Generate ULID
	newUlid, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating ULID: %v\n", err)
		os.Exit(1)
	}

	ulidStr := newUlid.String()
	args := flag.Args()

	// Determine output based on flags and arguments
	if len(args) == 0 {
		// No argument provided, just print ULID
		fmt.Println(ulidStr)
	} else {
		// Argument provided
		argStr := args[0]
		if *suffixFlag {
			// Append as suffix
			fmt.Println(ulidStr + argStr)
		} else {
			// Prepend as prefix (default)
			fmt.Println(argStr + ulidStr)
		}
	}
}

