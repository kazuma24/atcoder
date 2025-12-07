package main

import "fmt"

func main() {

	var a int //１円硬貨
	var n int //支払金額

	fmt.Scanf("%d\n%d", &n, &a)

	// fmt.Scanf("%d", &a)
	var s int //500円硬貨で割った残り

	s = n % 500
	if s <= a {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

//https://atcoder.jp/contests/abc088/tasks/abc088_a
