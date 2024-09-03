package mylib

import (
	"flag"
	"fmt"
)

func ParseFlagAndArgument() (uint, string, error) {
	var numOfLines uint
	flag.UintVar(&numOfLines, "n", 10, "number of lines")
	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return 0, "", fmt.Errorf("invalid file path, you have to specify a valid path to text file")
	}

	return numOfLines, filePath, nil
}
