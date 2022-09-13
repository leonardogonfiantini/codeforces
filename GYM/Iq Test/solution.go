package main

import ( 
	"fmt"
)


func main() {

	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		var x, y, z int
		fmt.Scan(&x, &y, &z)

		xmod := x % 2
		ymod := y % 2
		zmod := z % 2

		if xmod != ymod && xmod != zmod { 
			fmt.Print(3*i+1) 
			return
		}
		if ymod != xmod && ymod != zmod { 
			fmt.Print(3*i+2) 
			return
		}
		if zmod != xmod && zmod != ymod { 
			fmt.Print(3*i+3) 
			return
		}

	}

}