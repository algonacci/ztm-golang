//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name)
}

func hiThere() string {
	return "Hi there!"
}

func sumThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func returnAnyNumber() int {
	return 6
}

func returnAnyTwoNumbers() (int, int) {
	return 3, 5
}

func main() {
	greeting("Eric")

	fmt.Println(hiThere())

	fmt.Println("The total of three numbers are", sumThreeNumbers(7, 2, 1))

	anyNumber := returnAnyNumber()
	fmt.Println("Number from returnAnyNumber function is", anyNumber)

	firstNumber, secondNumber := returnAnyTwoNumbers()
	fmt.Println("Any two numbers are", firstNumber, "and", secondNumber)

	a, b := returnAnyTwoNumbers()
	calculateThreeNumbers := sumThreeNumbers(a, b, returnAnyNumber())
	fmt.Println("The calculation betweet three number is", calculateThreeNumbers)
}
