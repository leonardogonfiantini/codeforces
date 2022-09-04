package main

import (
	"fmt"
)

func main () {

	var candidates int
	var enemy int
	
	fmt.Scanln(&candidates, &enemy)

	can := make([]int, candidates)
	for i:=0; i<candidates; i++ {
		fmt.Scan(&can[i])
	}	
	
	//ordering player
	for i:=0; i <candidates-1; i++ {
		if can[i] < can[i+1] {
			temp := can[i]
			can[i] = can[i+1]
			can[i+1] = temp
		}
	}

	count := 0
	for i:=0; i<candidates; i++ {
		if can[i] > enemy { 
			count++ 
		} else {
			///can find solution///
		}
	}

}