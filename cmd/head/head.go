package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFlags() int {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines")
	flag.Parse()
	return n
}

func getArg() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("you have to specify an argument")
	}

	var argument string
	for i := 1; i < len(os.Args); i++ {
		if !strings.HasPrefix(os.Args[i], "-") {
			argument = os.Args[i]
			break
		}
	}
	return argument, nil

}

func readFromFile() []byte {
	file, err := getArg()
	check(err)

	data, err := os.ReadFile(file)
	check(err)

	return data
}

func print() {
	n := parseFlags()

	data := readFromFile()

	lines := 0
	for _, c := range data {
		if c == 10 {
			lines++
		}
		if lines < n {
			fmt.Print(string(c))
		} else {
			fmt.Println()
			break
		}

	}

}

func main() {
	print()
}
