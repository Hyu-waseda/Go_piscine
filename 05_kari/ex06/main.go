package main

import (
	"piscine"
	"fmt"
	//"strings"
)

func main() {
	fmt.Printf("%#v\n", piscine.Split("HelloHAhowHAareHAyou?", "HA"))
	//fmt.Printf("%#v\n", strings.Split("HelloHAhowHAareHAyou?", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HA", "HA"))
	//fmt.Printf("%#v\n", strings.Split("HA", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HAHAHA", "HA"))
	//fmt.Printf("%#v\n", strings.Split("HAHAHA", "HA"))
	fmt.Printf("%#v\n", piscine.Split("", "HA"))
	//fmt.Printf("%#v\n", strings.Split("", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HelloHAhowHAareHAyou?", ""))
	//fmt.Printf("%#v\n", strings.Split("HelloHAhowHAareHAyou?", ""))
	fmt.Printf("%#v\n", piscine.Split("", ""))
	//fmt.Printf("%#v\n", strings.Split("", ""))
}
