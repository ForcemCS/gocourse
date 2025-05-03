package basics

import "fmt"

var middleName string = "Cane"

func main() {

	middleName := "kui"
	fmt.Println(middleName)

	printname()

}

func printname() {
	firstName := "wu"
	fmt.Println(firstName)
	fmt.Println(middleName)
}
