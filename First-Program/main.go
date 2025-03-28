package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUsername() string {
	var name string
	fmt.Println("What is your name - ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumber() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter First Number: - ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number: - ")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	//Display Result
	fmt.Println("Hello , ", name)
	fmt.Println("Sum of the number is ", sum)
}

func printLast() {
	fmt.Println("Thank You, Good Bye")

}
func main() {

	printWelcomeMessage()
	name := getUsername()
	num1, num2 := getTwoNumber()
	sum := add(num1, num2)
	display(name, sum)
	printLast()

}
