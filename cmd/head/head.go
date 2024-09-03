package main

import (
	"flag"
	"fmt"
	"os"
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
	arg := flag.Arg(0)
	fmt.Println(arg)
	return arg, nil
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
