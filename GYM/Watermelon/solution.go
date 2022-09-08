package main

import ( "fmt" )

func main() {

	var input int
	fmt.Scan(&input)

	if input == 2 {
		fmt.Println("NO")
		return
	}

	if input % 2 != 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	
}