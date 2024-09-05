package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func parseFlagsAndArgument() (string, bool, bool, bool, error) {
	var lFlag, wFlag, cFlag bool

	flag.BoolVar(&lFlag, "l", false, "count lines flag")
	flag.BoolVar(&wFlag, "w", false, "count words flag")
	flag.BoolVar(&cFlag, "c", false, "count characters flag")

	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		return "", false, false, false, fmt.Errorf("invalid file path, you have to specify a valid path to text file")
	}

	return filePath, lFlag, wFlag, cFlag, nil
}

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
		_, err := file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		countLines(scanner)
	}

	if wFlag {
		_, err := file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		countWords(scanner)
	}

	if cFlag {
		_, err := file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		countCharacters(scanner)
	}

	if !lFlag && !wFlag && !cFlag {
		scanner := bufio.NewScanner(file)
		_, err := file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		countLines(scanner)

		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
		countWords(scanner)

		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
		countCharacters(scanner)
	}

	return nil
}

// another approach
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
	filePath, lFlag, wFlag, cFlag, err := parseFlagsAndArgument()
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
