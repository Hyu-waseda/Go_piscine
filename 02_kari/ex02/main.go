package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IterativePower(4, 3))
	fmt.Println(piscine.IterativePower(3, 4))
	fmt.Println(piscine.IterativePower(-4, 3))
	fmt.Println(piscine.IterativePower(-3, 4))
	fmt.Println(piscine.IterativePower(0, 3))
	fmt.Println(piscine.IterativePower(3, 0))
}
