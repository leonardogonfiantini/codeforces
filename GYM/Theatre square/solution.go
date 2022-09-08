package main

import (
	"fmt"
	"math"
)


func main() {

	var n, m, a float64
	fmt.Scan(&n, &m, &a)

	sol := (n*m) / (a*a)
	if sol - float64(math.Round(sol)) > 0 { sol+=2 }


	fmt.Println(int(sol))


}