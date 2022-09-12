package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)


	x,y,z := 0, 0, 0
	for i := 0; i < n; i++ {

		var xt, yt, zt int
		fmt.Scan(&xt, &yt, &zt)

		x += xt
		y += yt
		z += zt

	}

	if x == 0 && y == 0 && z == 0 { 
		fmt.Println("YES") 
	} else {
		fmt.Println("NO")
	}
	

	
	
}