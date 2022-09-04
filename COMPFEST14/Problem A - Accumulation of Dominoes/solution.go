package main

import (
	"fmt"
)


func main() {

	var row int
	var col int

	fmt.Scanln(&row, &col)

	count := (col-1) * row
	if col == 1 {
		count = row-1
	}

	fmt.Println(count)

}