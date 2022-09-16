package main

import (
	"fmt"
	"strconv"
	"bufio"
	"io"
	"os"
)

func main() {

	var n int
	fmt.Scan(&n)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	db := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		fmt.Fscan(r, &name)

		if db[name] == 0 {
			io.WriteString(w, "OK\n")
			db[name] = 1
		} else {
			io.WriteString(w, name+strconv.Itoa(db[name])+"\n")
			db[name] += 1
		}


	}

}