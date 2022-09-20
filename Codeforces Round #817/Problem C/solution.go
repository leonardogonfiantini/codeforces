package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanws(a ...interface{})           { fmt.Fscan(reader, a...) }


func main() {
	defer writer.Flush()

	var t int
	scanws(&t)

	for i := 0; i < t; i++ {
		
		var n int
		var a, b, c []string

		a,b,c = nil, nil, nil

		for j := 0; j < n; j++ {
			var x string
			scanws(&x)

			a = append(a, x)
		}

		for j := 0; j < n; j++ {
			var x string
			scanws(&x)

			b = append(b, x)
		}

		for j := 0; j < n; j++ {
			var x string
			scanws(&x)

			c = append(c, x)
		}


		var ascore, bscore, cscore int = 0, 0, 0

		for j := 0; j < n; j++ {
			
		}

	}

}