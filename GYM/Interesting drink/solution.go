package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)


	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	shops := make([]int, n)
	for i := 0; i < n; i++ {
		var shop int
		fmt.Fscan(r, &shop)

		shops[i] = shop
	}

	sort.Ints(shops)

	var d int
	fmt.Fscan(r, &d)

	for i := 0; i < d; i++ {
		var coin int
		fmt.Fscan(r, &coin)
		
		for j := n/2; j < n; j++ {
			if shops[j] > coin {
				j = j / 2
			}
		}
			
	}

//completare

}