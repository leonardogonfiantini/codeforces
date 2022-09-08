package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Scan(&n)

	strings := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		strings[i] = s
	}

	for i := 0; i < n; i++ {
		if len(strings[i]) > 10 {


			size := len(strings[i]) - 1

			fmt.Println(string(strings[i][0]) +
						strconv.Itoa(len(strings[i][1:size]))+
						string(strings[i][size]))


		} else {
			fmt.Println(strings[i])
		}


	}

}