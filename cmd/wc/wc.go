package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/internal/wc"
)

func parseFlagsAndArgument() (string, bool, bool, bool, error) {
	var lFlag, wFlag, cFlag bool

	flag.BoolVar(&lFlag, "l", false, "count lines flag")
	flag.BoolVar(&wFlag, "w", false, "count words flag")
	flag.BoolVar(&cFlag, "c", false, "count characters flag")

	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return "", false, false, false, fmt.Errorf("invalid file path, you have to specify a valid path to text file")
	}

	return filePath, lFlag, wFlag, cFlag, nil
}

func main() {
	filePath, lFlag, wFlag, cFlag, err := parseFlagsAndArgument()
	if err != nil {
		log.Fatal(err)
	}

	if err := wc.Count(filePath, lFlag, wFlag, cFlag); err != nil {
		log.Fatal(err)
	}

}
