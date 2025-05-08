package basics

import "fmt"

func main() {

	//var arrayName  [size]elementType
	//是不可变的

	var number [5]int
	fmt.Println(number)
	//数组和切片中索引从0开始
	number[4] = 20
	fmt.Println(number)

	var cc [5]string
	fmt.Println(cc)
	fruits := [4]string{"Appel", "Binana", "Orange", "Grapes"}
	fmt.Println(fruits)
	fmt.Println("Third Fruit: ", fruits[2])

	originalArray := [3]int{1, 2, 3}
	//发生的是值拷贝，不影响原来的数组
	copiedArray := originalArray

	copiedArray[0] = 100
	fmt.Println(originalArray[0])

	for i := 1; i < len(originalArray); i++ {
		fmt.Println(originalArray[i])
	}

	for i, v := range originalArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	//_ 表示丢弃返回的值
	for _, v := range copiedArray {
		fmt.Printf("Value: %d\n", v)
	}

	a, _ := someFunction()
	fmt.Println(a)

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	fmt.Println(a1 == a2)

	//定义一个3 * 3 的矩阵
	//matrix := [3][3]int{}
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	aa := [3]int{1, 2, 3}
	var bb *[3]int //类型是指针（指向长度3的整数数组）
	bb = &aa       //分配原始数组aa所在的确切地址
	bb[0] = 22
	fmt.Println(aa, *bb) //指针的零值是nil

}

func someFunction() (int, int) {
	return 1, 2
}
