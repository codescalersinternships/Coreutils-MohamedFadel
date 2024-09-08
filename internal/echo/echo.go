package echo

import (
	"fmt"
	"strings"
)

func Echo(trailingLineFlag bool, words []string) {
	fmt.Print(strings.Join(words, " "))

	if !trailingLineFlag {
		fmt.Println()
	}
}
