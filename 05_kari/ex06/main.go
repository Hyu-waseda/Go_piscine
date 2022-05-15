package main

import (
	"piscine"
	"fmt"
	"strings"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", strings.Split(s, "HA"))
	s = "HA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", strings.Split(s, "HA"))
	s = "HAHAHA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", strings.Split(s, "HA"))
	s = ""
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", strings.Split(s, "HA"))
	s = "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, ""))
	fmt.Printf("%#v\n", strings.Split(s, ""))
}
