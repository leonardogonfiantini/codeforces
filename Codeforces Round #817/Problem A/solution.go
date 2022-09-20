package main

import (
	"bufio"
	"fmt"
	"os"
	"bytes"
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

	var l int
	var s []byte



	for i := 0; i < t; i++ {
		
		l, s = 0, nil
		scanws(&l)
		scanws(&s)

		f := 0
		timur := []byte("Timur")

		if l != 5 {
			printf("\nNO")
			continue
		}

		for j := 0; j < l; j++ {
			
			index := bytes.IndexByte(timur, s[j])
			if  index >= 0 {
				timur[index] = 0
			} else {
				printf("\nNO")
				f = 1
				break
			}
		}

		if f == 0 {
			printf("\nYES")
		}
		
		

	}


}