package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanws(a ...interface{})           { fmt.Fscan(reader, a...) }

func main() {

	defer writer.Flush()
	var tc, n int
	var x, y, c []int
	

	scanws(&tc)
	for t := 0; t < tc; t++ {
		
		x, y, c = nil, nil, nil

		scanws(&n)

		for i := 0; i < n; i++ {
			var a int
			scanws(&a)

			x = append(x, a)
		}


		for i := 0; i < n; i++ {
			var a int
			scanws(&a)

			y = append(y, a)
		}

		for i := 0; i < n; i++ {
			p := y[i] - x[i]

			c = append(c, p)
		}

		sort.Ints(c)

		nn := n - 1
		idx := 0
		ans := 0

		for idx < nn {
			if c[idx] + c[nn] >= 0 {
				ans++
				nn--
			}
			idx++
		}

		printf("%d\n", ans)


	}


}	