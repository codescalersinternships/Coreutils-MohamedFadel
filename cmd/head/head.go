package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func parseLineFlag() int {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines")
	flag.Parse()
	return n
}

func getFilePathArg() string {
	arg := flag.Arg(0)
	return arg
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
	numOfLines := parseLineFlag()

	filePath := getFilePathArg()

	printFileLines(filePath, numOfLines)

}
