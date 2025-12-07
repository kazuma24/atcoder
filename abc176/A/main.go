package main

import (
	"fmt"
)

func main() {

	var N, X, T int
	fmt.Scan(&N, &X, &T)
	// 切り上げで必要な回数を計算
	rounds := (N + X - 1) / X
	// 必要な時間
	result := rounds * T
	fmt.Println(result)

}

//https://atcoder.jp/contests/abc176/tasks/abc176_a
