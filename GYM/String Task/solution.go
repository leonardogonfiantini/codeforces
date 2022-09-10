package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	fmt.Scan(&s)

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "y", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "u", "")
	s = strings.ReplaceAll(s, "i", "")


	for _, c := range s {
		fmt.Print(".",string(c))
	}

}