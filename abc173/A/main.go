package main

import (
	"fmt"
)

func main() {

	var N int

	fmt.Scanf("%d", &N)

	if N%1000 == 0 {
		fmt.Print(0)
	} else {
		fmt.Print(1000 - (N % 1000))
	}
}

//https://atcoder.jp/contests/abc173/tasks/abc173_a
