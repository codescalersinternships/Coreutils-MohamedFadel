package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func parseFlagAndArgument() (int, string, error) {
	var numOfLines int
	flag.IntVar(&numOfLines, "n", 10, "number of lines")
	flag.Parse()

	filePath := flag.Arg(0)

	var ErrInvalidFlag = fmt.Errorf("invalid flag, you have to specify a positive integer number of lines")
	var ErrInvalidFilePath = fmt.Errorf("invalid file path, you have to specify a valid path to text file")

	if numOfLines <= 0 {
		return 0, "", ErrInvalidFlag
	}

	if filePath == "" {
		return 0, "", ErrInvalidFilePath
	}

	return numOfLines, filePath, nil
}

func printFileLines(filePath string, numOfLines int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	lineCounter := 0
	for scanner.Scan() && lineCounter < numOfLines {
		line := scanner.Text()
		lineCounter++
		fmt.Println(line)
	}
}

func main() {
	numOfLines, filePath, err := parseFlagAndArgument()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printFileLines(filePath, numOfLines)

}
