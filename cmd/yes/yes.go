package main

import (
	"flag"
	"fmt"
)

func parseArguments() []string {
	flag.Parse()

	strings := flag.Args()

	return strings
}

func main() {
	strings := parseArguments()

	if len(strings) != 0 {
		for {
			for _, string := range strings {
				fmt.Print(string, " ")
			}
			fmt.Println()
		}
	}

	for {
		fmt.Println("y")
	}

}
