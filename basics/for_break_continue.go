package basics

import "fmt"

func main() {
	//满足 <=5,才循环执行{}中的内容
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for index, num := range arr {
		fmt.Printf("Index: %d , Value: %d\n", index, num) //%d 十进制整数  	%v通用占位符（值）
	}

	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			continue //continue 表示跳过这次循环的剩余部分，直接进入下一次循环。
		}
		fmt.Println("Odd Number:", i)

		if i == 5 {
			break //如果 i 等于 5，就执行 break，跳出整个循环。
		}
	}

	//这个的取值是根据最后一行星星的取值决定的
	rows := 5

	//最外层一共走了五次，第二层走了4 + 3 + 2 + 1
	for i := 1; i <= rows; i++ {
		// inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line
	}

	for i := range 6 {
		i++
		fmt.Print(i, "|")
	}
}
