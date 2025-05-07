package basics

import "fmt"

func main() {

	// 对expression进行多个判断
	//case 默认只执行匹配的代码块，不会自动"穿透"到下一个 case,短路执行
	// switch expression{
	// case value1:
	// 	block of code
	// ....
	// default:
	// 	block of code
	// }

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("frult is apple")
	case "binana":
		fmt.Println("fruit is binana")
	default:
		fmt.Println("Unknow Fruit")
	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("it is a weekday!")
	default:
		fmt.Println("is is a weekend!")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10 ")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Numner is more than 20 ")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Graeter than 1 ")
		fallthrough //强制执行下一个 case 的代码块，即使那个 case 的条件不成立。
	case num == 2:
		fmt.Println("num is 2 ")
	default:
		fmt.Println("Not two")
	}
}
