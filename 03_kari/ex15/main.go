package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
	toConcat = []string{"Hello", "", "", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
	toConcat = []string{"Hello"}
	fmt.Println(piscine.Join(toConcat, ":"))
	toConcat = []string{""}
	fmt.Println(piscine.Join(toConcat, ":"))
	toConcat = []string{}
	fmt.Println(piscine.Join(toConcat, ":"))
}
