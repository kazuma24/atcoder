package main

import "fmt"

func main() {

	var a int //歳
	var b int //13歳以上の金額

	fmt.Scanf("%d %d", &a, &b)

	if a >= 13 {
		fmt.Print(b)
	} else if 6 <= a {
		fmt.Print(b / 2)
	} else {
		fmt.Print(0)
	}
}

//https://atcoder.jp/contests/abc127/tasks/abc127_a
