package main

import (
	"bufio"
	"fmt"
	"os"
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
		scanws(&n)

		var s1 string
		var s2 string

		scanws(&s1)
		scanws(&s2)

		gb := "GB"

		f := 0

		for j := 0; j < len(s1); j++ {
			
			if (s1[j] == gb[0] || s1[j] == gb[1]) && (s2[j] == gb[0] || s2[j] == gb[1]) {

				continue 
			} else {
				if (s1[j] != s2[j]) {
					printf("NO\n")
					f = 1
					break
				}
			}
			 

		}
		
		if f == 0 {
			printf("YES\n")
		}


	}
}