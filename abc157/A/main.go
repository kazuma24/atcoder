package main

import (
	"fmt"
)

func main() {

	var N int

	fmt.Scanf("%d", &N)

	if N%2 == 0 {
		fmt.Print((N / 2))
	} else {
		fmt.Print((N / 2) + 1)
	}
}

//https://atcoder.jp/contests/abc157/tasks/abc157_a
