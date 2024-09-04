package main

import (
	"flag"
	"fmt"
)

func parseFlagsAndArguments() (bool, []string) {
	var trailingLineFlag bool

	flag.BoolVar(&trailingLineFlag, "n", false, "omit trailing line after printing to STDOUT")

	flag.Parse()

	strings := flag.Args()

	return trailingLineFlag, strings
}

func echo(trailingLineFlag bool, strings []string) {
	if trailingLineFlag {
		for _, string := range strings {
			fmt.Print(string, " ")
		}
	} else {
		for _, string := range strings {
			fmt.Print(string, " ")
		}
		fmt.Println()
	}
}

func main() {
	trailingLineFlag, strings := parseFlagsAndArguments()
	echo(trailingLineFlag, strings)
}
