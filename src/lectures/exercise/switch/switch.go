//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {
	switch age := 14; {
	case age == 0:
		fmt.Println("Newborn")
	case age < 4:
		fmt.Println("Toddler")
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}
