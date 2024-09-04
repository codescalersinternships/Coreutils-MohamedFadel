package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/codescalersinternships/Coreutils-MohamedFadel/internal"
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
	numOfLines, filePath, _, _, _, err := utils.ParseFlagsAndArgument()
	if err != nil {
		log.Fatal(err)
	}

	err = printFileLines(filePath, numOfLines)
	if err != nil {
		log.Fatal(err)
	}
}
