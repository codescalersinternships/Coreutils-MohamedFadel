package main

import (
	"fmt"
	"os"
)

func main() {
	for _, ev := range os.Environ() {
		fmt.Println(ev)
	}
}
