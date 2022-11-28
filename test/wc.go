package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "这是一段中文！！"
	var count int = 0
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
}
