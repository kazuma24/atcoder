package main

import "fmt"

func main() {

	var x uint
	var a uint
	var b uint

	fmt.Scanf("%d\n%d\n%d", &x, &a, &b)

	fmt.Print((x - a) % b)

}

//https://atcoder.jp/contests/abc087/tasks/abc087_a
