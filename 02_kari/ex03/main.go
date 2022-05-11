package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))
	fmt.Println(piscine.RecursivePower(3, 4))
	fmt.Println(piscine.RecursivePower(-4, 3))
	fmt.Println(piscine.RecursivePower(-3, 4))
	fmt.Println(piscine.RecursivePower(0, 3))
	fmt.Println(piscine.RecursivePower(3, 0))
}
