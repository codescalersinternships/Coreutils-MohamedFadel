package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/codescalersinternships/Coreutils-MohamedFadel/internal"
)

func countLines(scanner *bufio.Scanner) {
	var counter uint = 0

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		counter++
	}
	fmt.Println("Number of lines:", counter)

}

func countWords(scanner *bufio.Scanner) {
	var counter uint = 0

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counter++
	}
	fmt.Println("Number of words:", counter)

}

func countCharacters(scanner *bufio.Scanner) {
	var counter uint = 0

	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		counter++
	}
	fmt.Println("Number of characters:", counter)

}

func count(filePath string, lFlag, wFlag, cFlag bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open this file for reading: No such file or directory")
	}
	defer file.Close()

	if lFlag {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		countLines(scanner)
	}

	if wFlag {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		countWords(scanner)
	}

	if cFlag {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		countCharacters(scanner)
	}

	if !lFlag && !wFlag && !cFlag {
		scanner := bufio.NewScanner(file)
		file.Seek(0, 0)
		countLines(scanner)

		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
		countWords(scanner)

		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
		countCharacters(scanner)
	}

	return nil
}

func countUsingASCI(filePath string, lFlag, wFlag, cFlag bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("cannot open this file for reading: No such file or directory")
	}

	var lineCounter, wordCounter int = 0, 0
	var characterCounter int = len(data)

	for _, ch := range data {
		if ch == 10 {
			lineCounter++
		}
		if ch == 32 || ch == 10 {
			wordCounter++
		}
	}

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

func main() {
	_, filePath, lFlag, wFlag, cFlag, err := utils.ParseFlagsAndArgument()
	if err != nil {
		log.Fatal(err)
	}

	err = count(filePath, lFlag, wFlag, cFlag)
	if err != nil {
		log.Fatal(err)
	}
	err = countUsingASCI(filePath, lFlag, wFlag, cFlag)
	if err != nil {
		log.Fatal(err)
	}
}
