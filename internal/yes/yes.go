package yes

import (
	"fmt"
	"strings"
)

func Yes(words []string) {
	toBePrinted := "y"
	if len(words) != 0 {
		toBePrinted = strings.Join(words, " ")
	}

	for {
		fmt.Println(toBePrinted)
	}
}
