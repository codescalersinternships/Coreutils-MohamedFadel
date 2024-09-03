package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func parseFlagAndArgument() (int, string) {
	var numOfLines int
	flag.IntVar(&numOfLines, "n", 10, "number of lines")
	flag.Parse()

	filePath := flag.Arg(0)

	return numOfLines, filePath
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
	numOfLines, filePath := parseFlagAndArgument()
	printFileLines(filePath, numOfLines)

}
