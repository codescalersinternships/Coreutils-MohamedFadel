package head

import (
	"bufio"
	"fmt"
	"os"
)

func Head(filePath string, numOfLines uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s: cannot open this file for reading: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() && counter < int(numOfLines) {
		line := scanner.Text()
		counter++
		fmt.Println(line)
	}
	return nil
}
