package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "al"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
