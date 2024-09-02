package main

import (
	"errors"
	"os"
)

func getArg() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("you have to specify an argument")
	}
	arg := os.Args[1]
	return arg, nil

}

func main() {
	arg, error := getArg()
}
