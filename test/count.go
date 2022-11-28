package main

import "fmt"

func main() {
	str := "asc中文개를１９あるノン"
	for i, ch := range str {
		fmt.Printf("%02d: %c\n", i, ch)
	}
}
