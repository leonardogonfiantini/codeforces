package main

import (
	"fmt"
	"sort"
)

func main() {

	var l, s int
	fmt.Scan(&l, &s)

	lanterns := make([]float64, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&lanterns[i])
	}

	sort.Float64s(lanterns)

	var d float64 = lanterns[0]
	
	for i := 0; i < l-1; i++ {
		td := (lanterns[i+1] - lanterns[i]) / 2
		if td > d {
			d = td
		}
	}

	fd := float64(s) - lanterns[l-1]
	if fd > d {d = fd }
	
	fmt.Print(d)
}