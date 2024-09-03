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

	var lineCounter uint = 0
	for scanner.Scan() && lineCounter < numOfLines {
		line := scanner.Text()
		lineCounter++
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
