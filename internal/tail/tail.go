package tail

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func Tail(filePath string, numOfLines uint) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Wrapf(err, "%s: cannot open this file for reading", filePath)
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if uint(len(lines)) > numOfLines {
		lines = lines[len(lines)-int(numOfLines):]
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
