package main

import (
	"flag"
)

func parseFlagsAndArguments() (bool, []string) {
	var trailingLineFlag bool

	flag.BoolVar(&trailingLineFlag, "n", false, "omit trailing line after printing to STDOUT")

	flag.Parse()

	strings := flag.Args()

	return trailingLineFlag, strings
}

func main() {
	trailingLineFlag, strings := parseFlagsAndArguments()

}
