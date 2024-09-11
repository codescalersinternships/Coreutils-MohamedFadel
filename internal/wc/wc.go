package wc

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func Count(filePath string, lFlag, wFlag, cFlag bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Wrapf(err, "cannot open this file for reading: %s", filePath)
	}

	characterCounter := len(data)
	lineCounter := len(strings.Split(string(data), "\n")) - 1
	wordCounter := len(strings.Fields(string(data)))

	if !lFlag && !wFlag && !cFlag {
		lFlag = true
		wFlag = true
		cFlag = true
	}

	if lFlag {
		fmt.Print(lineCounter, " ")
	}
	if wFlag {
		fmt.Print(wordCounter, " ")
	}
	if cFlag {
		fmt.Print(characterCounter, " ")
	}
	fmt.Println(filePath)

	return nil
}
