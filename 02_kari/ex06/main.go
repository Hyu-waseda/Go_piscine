package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(-2, piscine.IsPrime(-2))
	fmt.Println(-1, piscine.IsPrime(-1))
	fmt.Println(0, piscine.IsPrime(0))
	fmt.Println(1, piscine.IsPrime(1))
	fmt.Println(2, piscine.IsPrime(2))
	fmt.Println(3, piscine.IsPrime(3))
	fmt.Println(4, piscine.IsPrime(4))
	fmt.Println(5, piscine.IsPrime(5))
	fmt.Println(6, piscine.IsPrime(6))
	fmt.Println(7, piscine.IsPrime(7))
	fmt.Println(8, piscine.IsPrime(8))
	fmt.Println(9, piscine.IsPrime(9))
	fmt.Println(10, piscine.IsPrime(10))
	fmt.Println(10, piscine.IsPrime(57))
}
