package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

func cat(lineNumberingFlag bool, filesPaths []string) error {
	var lineCounter uint = 1

	for _, path := range filesPaths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("%s: cannot open this file for reading: No such file or directory", path)
		}

		scanner := bufio.NewScanner(file)

		if lineNumberingFlag {
			for scanner.Scan() {
				fmt.Println(lineCounter, "	", scanner.Text())
				lineCounter++
			}
		} else {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}

	return nil
}

func main() {
	lineNumberingFlag, filesPaths, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}
	err = cat(lineNumberingFlag, filesPaths)
	if err != nil {
		log.Fatal(err)
	}

}
