package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("hello from vim")

	mp := make(map[string]string)

	mp["apple"] = "223e"
	fmt.Println(mp)
	// arr:=[]int{1,2,3,4}
	m := mp["apple"]
	fmt.Println(len(m))
}
