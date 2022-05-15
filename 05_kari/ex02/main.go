package main

import (
	"piscine"
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
    fmt.Println(piscine.ConcatParams(test))
	fmt.Println("------------------")
	test = []string{"Hello", "", "", "you?"}
    fmt.Println(piscine.ConcatParams(test))
	fmt.Println("------------------")
	test = []string{""}
    fmt.Println(piscine.ConcatParams(test))
	fmt.Println("------------------")
	test = []string{}
    fmt.Println(piscine.ConcatParams(test))
}
