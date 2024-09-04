package main

import (
	"flag"
	"fmt"
	"log"
)

func parseFlagsAndArguments() (uint, string, error) {
	var depthLevelFlag uint

	flag.UintVar(&depthLevelFlag, "L", 1, "Max display depth of the directory tree")

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		return 1, "", fmt.Errorf("invalid path, you have to specify a valid path to a directory")
	}

	return depthLevelFlag, path, nil
}
func main() {
	depthLevelFlag, path, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}
}
