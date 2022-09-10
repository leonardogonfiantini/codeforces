package main

import (
	"fmt"
	"math"
)


func main() {

	var n, m, a float64
	fmt.Scan(&n, &m, &a)

	sol := math.Ceil(m/a) * math.Ceil(n/a)

	fmt.Println(int(sol))


}