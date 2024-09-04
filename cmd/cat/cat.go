package main

import (
	"flag"
	"fmt"
	"log"
)

func parseFlagsAndArguments() (bool, []string, error) {
	var lineNumberingFlag bool

	flag.BoolVar(&lineNumberingFlag, "n", false, "print content with line numbers")

	flag.Parse()

	filesPaths := flag.Args()
	if len(filesPaths) == 0 {
		return false, []string{}, fmt.Errorf("invalid file(s) path(s), you have to specify at least one valid path to text file")
	}

	return lineNumberingFlag, filesPaths, nil
}

func main() {
	_, _, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}

}
