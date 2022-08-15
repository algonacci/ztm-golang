package main

import "fmt"

func main() {
	var myName = "Eric"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	userName := "Admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "and other is", other)

	part2, other := 2, 6
	fmt.Println("part 2 is", part2, "and other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson name is", lessonName, "and type is", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
