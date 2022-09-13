package main

import ( 
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	var c1, c2, c3, c4 = 0, 0, 0, 0


	for i := 0; i < n; i++ {

		var x int 
		fmt.Scan(&x)

		switch x {
			case 1: c1++
			case 2: c2++
			case 3: c3++
			case 4: c4++
		}

	}


	var ans int = c4

	if c3 <= c1 {
		ans += c3
		c1 -= c3
		c3 = 0
	} else {
		ans += c1
		c3 -= c1
		c1 = 0
	}

	ans += c3

	for c1 >= 2 && c2 >= 1 {
		c2--
		c1 -= 2
		ans++
	}

	if c2 % 2 == 0 {
		c2 /= 2
		ans += c2
		c2 = 0
	} else {
		ans += (c2-1) / 2
		c2 -= c2-1
	}

	if c2 >= 1 && c1 >= 1 {
		c1--
		c2--
		ans++
	}

	ans += c2

	if c1 > 0 {
		ans += int(math.Ceil(float64(c1)/4))
	}

	fmt.Println(ans)



}