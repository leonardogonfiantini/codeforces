package main


import ( "fmt" )

func main() {

	var n int
	fmt.Scan(&n)

	counter := 0
	for i := 0; i < n; i++ {
		var p, v, t int
		fmt.Scan(&p, &v, &t)

		c := 0
		if p == 1 { c++ }
		if v == 1 { c++ }
		if t == 1 { c++ }

		if c > 1 { counter++ }

	}

	fmt.Println(counter)


}