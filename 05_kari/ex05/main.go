package main

import (
	"piscine"
	"fmt"
)

func main() {
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
	result = piscine.ConvertBase("43", "0123456789", "01")
	fmt.Println(result)
}
