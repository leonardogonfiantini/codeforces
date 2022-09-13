package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	dbmap := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)

		if dbmap[name] > 0 {
			fmt.Printf("%s%d\n", name, dbmap[name])
		} else {
			fmt.Println("OK")
		}

		dbmap[name] += 1

	}

}