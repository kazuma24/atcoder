package main

import (
	"fmt"
)

func main() {

	var n int
	var a int
	var b int

	fmt.Scanf("%d %d %d", &n, &a, &b)

	var r int
	for n > 0 {
		rr := sumDigit(n)
		if a <= rr && rr <= b {
			r += n
		}
		n--
	}
	fmt.Println(r)

}

func sumDigit(n int) int {
	var sum int
	for n > 0 {
		//整数nを10で割った余りを取得（末尾一桁）
		sum += n % 10
		//整数nを10で割った商をnとする
		n = n / 10
	}
	// fmt.Println(sum)
	return sum
}

//https://atcoder.jp/contests/abc083/tasks/abc083_b
