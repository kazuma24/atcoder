package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Scanf("%d %d", &a, &b)

	if (b % a) == 0 {
		fmt.Print(a + b)
	} else {
		fmt.Print(b - a)
	}
}

//https://atcoder.jp/contests/abc118/tasks/abc118_a
