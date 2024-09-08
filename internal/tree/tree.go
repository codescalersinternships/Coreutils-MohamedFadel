package tree

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func tab(tabs uint) {
	for i := 0; i < int(tabs); i++ {
		fmt.Print("│   ")
	}
}

func Tree(path string, depthLevelFlag, numOfTabs uint, dirCount, fileCount *uint) error {
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

			if err := Tree(newPath, depthLevelFlag-1, numOfTabs+1, dirCount, fileCount); err != nil {
				log.Fatal(err)
			}

		}

	}

	return nil
}
