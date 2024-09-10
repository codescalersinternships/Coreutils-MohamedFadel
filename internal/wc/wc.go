package wc

import (
	"fmt"
	"os"
	"strings"
)

func Count(filePath string, lFlag, wFlag, cFlag bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%s: cannot open this file for reading: %w", filePath, err)
	}

	characterCounter := len(data)
	lineCounter := len(strings.Split(string(data), "\n")) - 1
	wordCounter := len(strings.Fields(string(data)))

	if !lFlag && !wFlag && !cFlag {
		lFlag = true
		wFlag = true
		cFlag = true
	}

	if lFlag {
		fmt.Print(lineCounter, " ")
	}
	if wFlag {
		fmt.Print(wordCounter, " ")
	}
	if cFlag {
		fmt.Print(characterCounter, " ")
	}
	fmt.Println(filePath)

	return nil
}
