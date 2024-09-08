package cat

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(lineNumberingFlag bool, filesPaths []string) error {
	var lineCounter uint = 1

	for _, path := range filesPaths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("%s: cannot open this file for reading: No such file or directory", path)
		}

		scanner := bufio.NewScanner(file)

		if lineNumberingFlag {
			for scanner.Scan() {
				fmt.Println(lineCounter, "  ", scanner.Text())
				lineCounter++
			}
		} else {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	}

	return nil
}
