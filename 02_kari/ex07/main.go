package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(-2, piscine.FindNextPrime(-2))
	fmt.Println(-1, piscine.FindNextPrime(-1))
	fmt.Println(0, piscine.FindNextPrime(0))
	fmt.Println(1, piscine.FindNextPrime(1))
	fmt.Println(2, piscine.FindNextPrime(2))
	fmt.Println(3, piscine.FindNextPrime(3))
	fmt.Println(4, piscine.FindNextPrime(4))
	fmt.Println(5, piscine.FindNextPrime(5))
	fmt.Println(6, piscine.FindNextPrime(6))
	fmt.Println(7, piscine.FindNextPrime(7))
	fmt.Println(8, piscine.FindNextPrime(8))
	fmt.Println(9, piscine.FindNextPrime(9))
	fmt.Println(10, piscine.FindNextPrime(10))
	fmt.Println(10, piscine.FindNextPrime(57))
}
