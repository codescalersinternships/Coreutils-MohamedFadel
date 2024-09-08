package main

import (
	"flag"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/internal/echo"
)

func parseFlagsAndArguments() (bool, []string) {
	var trailingLineFlag bool

	flag.BoolVar(&trailingLineFlag, "n", false, "omit trailing line after printing to STDOUT")

	flag.Parse()

	words := flag.Args()

	return trailingLineFlag, words
}

func main() {
	trailingLineFlag, words := parseFlagsAndArguments()
	echo.Echo(trailingLineFlag, words)
}
