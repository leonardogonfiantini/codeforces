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
	var x, y []int
	var cst []int
 
	scanws(&tc)
	for t := 1; t <= tc; t++ {
		scanws(&n)
		x = nil
		y = nil
		cst = nil
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
			c := y[i] - x[i]
			cst = append(cst, c)
		}
		sort.Slice(cst, func(i, j int) bool {
			return cst[i] < cst[j]
		})
		
		nn := n - 1
		idx := 0
		ans := 0
		for idx < nn {
			if cst[idx]+cst[nn] >= 0 {
				ans++
				nn--
			}
			idx++
		}
		printf("%d\n", ans)
	}
 
}