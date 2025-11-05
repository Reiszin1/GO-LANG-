package main

import "fmt"

var (
	number1 = 2
	number2 = 6
)

func main() {

	// var name string = "John"
	// fmt.Println(name)

	name := "John"
	age := 20
	dev := false

	fmt.Println("Hello,", name, "you are", age, "years old")
	fmt.Println("--------------------------------")
	if dev {
		fmt.Println("You are a developer")
	} else {
		fmt.Println("You are not a developer")
	}
	fmt.Println("--------------------------------")

	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}
	fmt.Println("--------------------------------")
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println("i is", i)
		fmt.Println("the sum of the number", sum, "and", i, "is")
		sum += i
		fmt.Println(sum)
		fmt.Println("--------------------------------")
	}

	fruits := []string{"apple", "banana", "cherry"}
	numbers := []int{1, 2, 3, 4, 5}

	for i, fruit := range fruits {
		fmt.Println("index", i, "fruit", fruit)
		fmt.Println("--------------------------------")
	}
	for i, number := range numbers {
		fmt.Println("index", i, "number", number)
		fmt.Println("--------------------------------")
	}

	sum, err := somar(number1, number2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the sum of", number1, "and", number2, "is", sum)
	}
	fmt.Println("--------------------------------")

	subtraction, err := subtrair(number1, number2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the subtraction of", number1, "and", number2, "is", subtraction)
	}
	fmt.Println("--------------------------------")
}
