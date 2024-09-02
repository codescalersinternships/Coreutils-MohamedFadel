package main

import (
	"errors"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getArg() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("you have to specify an argument")
	}

	var argument string
	for i := 1; i < len(os.Args); i++ {
		if !strings.HasPrefix(os.Args[i], "-") {
			argument = os.Args[i]
			break
		}
	}
	return argument, nil

}

func main() {

}
