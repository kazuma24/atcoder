package main

import "fmt"

func main() {

	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	data := []int{}
	data = append(data, (a + b))
	data = append(data, (a - b))
	data = append(data, (a * b))

	r := data[0]
	for _, v := range data[1:] {
		if r <= v {
			r = v
		}
	}

	fmt.Print(r)
}

// https://atcoder.jp/contests/abc137/tasks/abc137_a
