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

	var lineCounter uint = 0
	for scanner.Scan() && lineCounter < numOfLines {
		line := scanner.Text()
		lineCounter++
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
