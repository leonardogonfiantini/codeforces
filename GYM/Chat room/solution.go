package main

import (
	"fmt"
)

const hello string = "hello"

func main() {
	var s []byte
	fmt.Scan(&s)

	p := 0
	for _, c := range s {

		if c == hello[p] {
			p++
		}

		if p == 5 {
			fmt.Println("YES")
			return
		}

	}

	fmt.Println("NO")
	
}
