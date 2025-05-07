package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//time.Now().UnixNano()：获取当前时间的纳秒数，作为种子值（seed）。
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("欢迎来到猜数字游戏!")
	fmt.Println("选这介于1到100的数字")
	fmt.Println("你可以猜到我是什么?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		//输入的内容存储在这个地址代表的变量中，当我们访问变量的时候，就会得到正确的值
		//如果只是guess，则输入的值是存储在变量的副本中，而不是原始变量
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("恭喜，你猜对了！")
			break
		} else if guess < target {
			fmt.Println("你猜小了！")
		} else {
			fmt.Println("你猜大了！")
		}
	}
}
