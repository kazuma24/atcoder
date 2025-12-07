package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var T int

	fmt.Scanf("%d %d %d", &a, &b, &T)

	r := (T / a) * b
	fmt.Print(int(r))
}

//https://atcoder.jp/contests/abc125/tasks/abc125_a
