package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func binarySearch(arr []int, start int, end int, coin int) int {

	if start > end {
		return 0
	}

	middle := (end + start) / 2

	if (arr[middle] <= coin && arr[middle+1] > coin) || arr[middle] == coin {
		return middle+1
	} else {

		if arr[middle] > coin {
			return binarySearch(arr, start, middle-1, coin)
		} else {
			return binarySearch(arr, middle+1, end, coin)
		}

	}

}

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
		
		count := binarySearch(shops, 0, n-1, coin)

		io.WriteString(w, strconv.Itoa(count)+"\n")
			
	}


	//completare
}