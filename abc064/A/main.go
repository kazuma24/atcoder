package main

import (
	"fmt"
)

func main() {

	var r int
	var g int
	var b int

	fmt.Scanf("%d %d %d", &r, &g, &b)

	num := r*100 + g*10 + b

	if num%4 == 0 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

//https://atcoder.jp/contests/abc087/tasks/abc087_a
