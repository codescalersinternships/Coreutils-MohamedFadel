package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func parseFlags() int {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines")
	flag.Parse()
	return n
}

func getArg() string {
	arg := flag.Arg(0)
	return arg
}

func printFileLines(fileName string, nLines int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	lineCounter := 0
	for scanner.Scan() && lineCounter < nLines {
		line := scanner.Text()
		lineCounter++
		fmt.Println(line)
	}
}

func main() {
	nLines := parseFlags()

	file := getArg()

	printFileLines(file, nLines)

}
