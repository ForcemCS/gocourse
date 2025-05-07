// for不仅可以用于传统的计数循环，还可以作为while循环，根据条件重复执行代码块
package basics

import "fmt"

func main() {

	//这是一个无限循环
	// for  {
	// 	fmt.Println("Hello")
	// }

	// i := 1
	// for i <= 5 {

	// 	fmt.Println(i)
	// 	i++

	// }

	sum := 0
	for {
		sum += 10
		fmt.Println("Sum: ", sum)

		if sum >= 50 {
			break
		}
	}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number: ", num)
		num++

	}
}
