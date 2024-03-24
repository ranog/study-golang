package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for i, arg := range os.Args[1:] {
		fmt.Println(fmt.Sprint(i) + sep + arg)
	}
	fmt.Println(s)
}
