package main

import (
	"fmt"
)

func main() {

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		
		var n, k int
		fmt.Scan(&n, &k)

		if n <= 3 {
			c := 0
			for j := 0; j < n; j++ {
				var x int
				fmt.Scan(&x)
				c += x
			}	

			fmt.Println(c)
			continue
		}

		arr := make([]int, n)
		for j := 0; j < n; j++ {
			var x int
			fmt.Scan(&x)

			arr[j] = x
		}	

		tk := k
		for j := 0; j < n-k; j++ {
			if arr[j] > arr[j+k] && tk != 0 {
				temp := arr[j]
				arr[j] = arr[j+k]
				arr[j+k] = temp
			}
		}


		score := 0
		for j := 0; j <= n-k; j++ {

			p := 0
			for l := j; l < k+j; l++ {
				p += arr[l]
			}

			if p > score {
				score = p
			}
		}

		fmt.Println(score)

	}	

}