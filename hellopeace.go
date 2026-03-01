package main

import "fmt"

func main() {
	fmt.Println("Hello, Peace!")

	//	Declare variables with explicit type and inferred type
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//	Declare variables with zero values
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//	Declare variables with var and short variable declaration
	var student1 string
	student1 = "John"
	fmt.Println(student1)
}
