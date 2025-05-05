package main

import "fmt"

var a, b = 10, 3

func main() {
	//整数的除法知会保留整数部分
	result := a / b
	fmt.Println(result)

	//结果存储在一个浮点变量中
	const p float64 = 22 / 7.0
	fmt.Println(p)
	fmt.Printf("%.4f\n", p)
	//int64的最大值是2^63  -1 ,最小是   -2^63  -1
	//uint64的最大值是2^64 -1
	var maxInt int64 = 9223372036854775807
	cc := maxInt + 1
	fmt.Println(cc)

	//
}
