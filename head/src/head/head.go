package main

import (
	"errors"
	"fmt"
	"os"
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
	arg := os.Args[1]
	return arg, nil

}

func main() {
	arg, error := getArg()
	check(error)
	fmt.Println(arg)

}
