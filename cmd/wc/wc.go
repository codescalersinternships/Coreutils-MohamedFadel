package main

import (
	"bufio"
	"fmt"
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

func main() {

}
