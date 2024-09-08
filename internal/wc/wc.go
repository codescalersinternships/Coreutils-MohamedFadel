package wc

import (
	"fmt"
	"os"
	"strings"
)

func Count(filePath string, lFlag, wFlag, cFlag bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("cannot open this file for reading: No such file or directory")
	}

	var characterCounter int = len(data)
	var lineCounter = len(strings.Split(string(data), "\n")) - 1
	var wordCounter = len(strings.Split(string(data), " ")) + lineCounter - 1

	if lFlag {
		fmt.Println("Number of lines:", lineCounter)
	}
	if wFlag {
		fmt.Println("Number of words:", wordCounter)
	}
	if cFlag {
		fmt.Println("Number of characters:", characterCounter)
	}

	if !lFlag && !wFlag && !cFlag {
		fmt.Println("Number of lines:", lineCounter)
		fmt.Println("Number of words:", wordCounter)
		fmt.Println("Number of characters:", characterCounter)

	}

	return nil
}
