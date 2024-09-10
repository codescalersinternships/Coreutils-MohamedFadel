package tree

import (
	"fmt"
	"os"
	"strings"
)

func tab(tabs uint) {
	for i := 0; i < int(tabs); i++ {
		fmt.Print("│   ")
	}
}

func Tree(startPath, currentPath string, depthLevelFlag, numOfTabs uint) (uint, uint, error) {
	var fileCount, dirCount uint = 0, 0

	if startPath == currentPath {
		fmt.Println(startPath)
	}

	contents, err := os.ReadDir(currentPath)
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %w", currentPath, err)
	}

	for _, content := range contents {
		if strings.HasPrefix(content.Name(), ".") {
			continue
		}

		tab(numOfTabs)
		fmt.Println("│──", content.Name())

		if content.IsDir() {
			dirCount++

			if depthLevelFlag > 1 {
				newPath := currentPath + "/" + content.Name()
				subDirCount, subFileCount, err := Tree(startPath, newPath, depthLevelFlag-1, numOfTabs+1)

				if err != nil {
					return 0, 0, err
				}

				dirCount += subDirCount
				fileCount += subFileCount
			}
		} else {
			fileCount++
		}
	}

	if currentPath == startPath {
		fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
	}

	return dirCount, fileCount, nil
}
