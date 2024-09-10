package cat

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(lineNumberingFlag bool, filesPaths []string) error {
	counter := 1

	for _, path := range filesPaths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("%s: cannot open this file for reading: %w", path, err)
		}

		scanner := bufio.NewScanner(file)

		if lineNumberingFlag {
			for scanner.Scan() {
				fmt.Println(counter, "  ", scanner.Text())
				counter++
			}
		} else {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}

	return nil
}
