package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(-2, piscine.RecursiveFactorial(-2))
	fmt.Println(-1, piscine.RecursiveFactorial(-1))
	fmt.Println(0, piscine.RecursiveFactorial(0))
	fmt.Println(1, piscine.RecursiveFactorial(1))
	fmt.Println(2, piscine.RecursiveFactorial(2))
	fmt.Println(3, piscine.RecursiveFactorial(3))
	fmt.Println(4, piscine.RecursiveFactorial(4))
	fmt.Println(5, piscine.RecursiveFactorial(5))
	fmt.Println(6, piscine.RecursiveFactorial(6))
	fmt.Println(7, piscine.RecursiveFactorial(7))
	fmt.Println(8, piscine.RecursiveFactorial(8))
	fmt.Println(9, piscine.RecursiveFactorial(9))
	fmt.Println(10, piscine.RecursiveFactorial(10))
	fmt.Println(20, piscine.RecursiveFactorial(20))
	fmt.Println(21, piscine.RecursiveFactorial(21))
}
