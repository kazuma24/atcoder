package main

import (
	"fmt"
)

func main() {

	var k int
	var a int
	var b int

	fmt.Scanf("%d\n%d %d", &k, &a, &b)

	// kk := 1
	// var r bool
	// for {
	// 	var v = k * kk

	// 	if v > b {
	// 		break
	// 	}

	// 	if a <= v && v <= b {
	// 		r = true
	// 		break
	// 	}

	// 	kk += 1
	// }

	// if r {
	// 	fmt.Print("OK")
	// } else {
	// 	fmt.Print("NG")
	// }

	var r bool
	for i := a; i <= b; i++ {
		if i%k == 0 {
			r = true
		}
	}
	if r {
		fmt.Print("OK")
	} else {
		fmt.Print("NG")
	}
}

//https://atcoder.jp/contests/abc165/tasks/abc165_a
