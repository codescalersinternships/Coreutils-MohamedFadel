package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

func printFileLines(filePath string, numOfLines uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open this file for reading: No such file or directory")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0, numOfLines)

	for scanner.Scan() {
		line := scanner.Text()

		if uint(len(lines)) == numOfLines {
			lines = lines[1:]
		}

		lines = append(lines, line)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func main() {
	numOfLines, filePath, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}

	err = printFileLines(filePath, numOfLines)
	if err != nil {
		log.Fatal(err)
	}
}
