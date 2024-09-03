package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codescalersinternships/Coreutils-MohamedFadel/mylib"
)

func printFileLines(filePath string, numOfLines uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open this file for reading: No such file or directory")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0, numOfLines)

	for scanner.Scan() {
		line := scanner.Text()

		if uint(len(lines)) == numOfLines {
			lines = lines[1:]
		}

		lines = append(lines, line)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func main() {
	numOfLines, filePath, err := mylib.ParseFlagAndArgument()
	mylib.HandleError(err)
	err = printFileLines(filePath, numOfLines)
	mylib.HandleError(err)
}
