package cat

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func Cat(lineNumberingFlag bool, filesPaths []string) error {
	counter := 1

	for _, path := range filesPaths {
		file, err := os.Open(path)
		if err != nil {
			return errors.Wrapf(err, "%s: cannot open this file for reading", path)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			var prefix string

			if lineNumberingFlag {
				prefix = fmt.Sprintf("%d ", counter)
				counter++
			}

			fmt.Println(prefix, scanner.Text())
		}
	}

	return nil
}
