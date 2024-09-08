package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/internal/tail"
)

func parseFlagsAndArguments() (uint, string, error) {
	var numOfLines uint

	flag.UintVar(&numOfLines, "n", 10, "number of lines")

	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return 0, "", fmt.Errorf("invalid file path, you have to specify a valid path to text file")
	}

	return numOfLines, filePath, nil
}

func main() {
	numOfLines, filePath, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}

	if err = tail.Tail(filePath, numOfLines); err != nil {
		log.Fatal(err)
	}

}
