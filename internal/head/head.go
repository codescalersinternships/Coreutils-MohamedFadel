package head

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func Head(filePath string, numOfLines uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrapf(err, "%s: cannot open this file for reading: ", filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counter uint = 0
	for scanner.Scan() && counter < numOfLines {
		line := scanner.Text()
		counter++
		fmt.Println(line)
	}
	return nil
}
