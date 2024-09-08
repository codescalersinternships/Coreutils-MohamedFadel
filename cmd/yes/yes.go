package main

import (
	"flag"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/internal/yes"
)

func parseArguments() []string {
	flag.Parse()
	return flag.Args()
}

func main() {
	words := parseArguments()

	yes.Yes(words)

}
