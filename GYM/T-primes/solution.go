package main 


import (
	"fmt"
	"math"
)


func isTPrime(x float64) bool {

	for i := 2; i <= int(math.Sqrt(x))+1; i++ {
		if int(x) % i == 0 {
			return false
		}	

	}

	return true

}

func main() {

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		
		var x float64
		fmt.Scan(&x)

		if isTPrime(x) { 
			fmt.Println("YES") 
		} else {
			fmt.Println("NO")
		}

	}

}