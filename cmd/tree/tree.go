package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func tab(tabs uint) {
	for i := 0; i < int(tabs); i++ {
		fmt.Print("│   ")
	}
}

func tree(path string, depthLevelFlag, numOfTabs uint, dirCount, fileCount *uint) error {
	contents, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("%s: [error opening dir]", path)
	}

	for _, content := range contents {
		if strings.HasPrefix(content.Name(), ".") {
			continue
		}

		if content.IsDir() {
			*dirCount++
		} else {
			*fileCount++
		}

		tab(numOfTabs)
		fmt.Println("│──", content.Name())

		if content.IsDir() && depthLevelFlag > 1 {
			newPath := path + "/" + content.Name()
			numOfTabs++
			depthLevelFlag--

			err := tree(newPath, depthLevelFlag, numOfTabs, dirCount, fileCount)
			if err != nil {
				log.Fatal(err)
			}

			numOfTabs--
			depthLevelFlag++
		}

	}

	return nil
}
func main() {
	depthLevelFlag, path, err := parseFlagsAndArguments()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)

	var fileCount, dirCount uint = 0, 0

	err = tree(path, depthLevelFlag, 0, &dirCount, &fileCount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(dirCount, "directories,", fileCount, "files")

}
