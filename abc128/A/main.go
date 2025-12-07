package main

import (
	"fmt"
)

func main() {

	var a int
	var p int

	fmt.Scanf("%d %d", &a, &p)

	k := (a * 3) + p
	fmt.Print(int(k / 2))
}

//https://atcoder.jp/contests/abc128/tasks/abc128_a
