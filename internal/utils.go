package utils

import (
	"flag"
	"fmt"
	"log"
)

func ParseFlagAndArgument() (uint, string, bool, bool, bool, error) {
	var numOfLines uint
	var lFlag, wFlag, cFlag bool

	flag.UintVar(&numOfLines, "n", 10, "number of lines")
	flag.BoolVar(&lFlag, "l", false, "count lines flag")
	flag.BoolVar(&wFlag, "w", false, "count words flag")
	flag.BoolVar(&cFlag, "c", false, "count characters flag")
	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return 0, "", false, false, false, fmt.Errorf("invalid file path, you have to specify a valid path to text file")
	}

	return numOfLines, filePath, lFlag, wFlag, cFlag, nil
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
