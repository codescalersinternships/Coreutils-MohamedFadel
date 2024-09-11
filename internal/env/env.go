package env

import (
	"fmt"
	"os"
)

func Env() {
	for _, ev := range os.Environ() {
		fmt.Println(ev)
	}
}
