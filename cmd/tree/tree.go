package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/internal/tree"
)

func parseFlagsAndArguments() (uint, string, error) {
	var depthLevelFlag uint
	flag.UintVar(&depthLevelFlag, "L", 1, "Max display depth of the directory tree")
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return 1, "", fmt.Errorf("invalid path, you have to specify a valid path to a directory")
		}
		path = currentDir
	}

	return depthLevelFlag, path, nil
}

func main() {
	depthLevelFlag, path, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = tree.Tree(path, path, depthLevelFlag, 0)
	if err != nil {
		log.Fatal(err)
	}
}
