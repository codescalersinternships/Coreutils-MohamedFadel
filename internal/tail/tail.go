package tail

import (
	"bufio"
	"fmt"
	"os"
)

func Tail(filePath string, numOfLines uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s: cannot open this file for reading: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if uint(len(lines)) > numOfLines {
		lines = lines[len(lines)-int(numOfLines):]
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
