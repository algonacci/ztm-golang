package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	GraduationYear int
}

func greeting(name string, student Student) {
	fmt.Println("Hello", name, "Welcome to the class of", student.GraduationYear)
}

func main() {
	var student = Student{Person{"Eric", 21}, 2023}
	fmt.Println(student.Age)
	fmt.Println(student.GraduationYear)
	greeting(student.Name, student)
}
