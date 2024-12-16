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
	//--Requirements:
	//* Use a `switch` statement to print the following:
	switch age := 45; {
	//  - "newborn" when age is 0
	case age == 0:
		fmt.Println("newborn")
	//  - "toddler" when age is 1, 2, or 3
	case age < 3:
		fmt.Println("toddler")
	//  - "child" when age is 4 through 12
	case age < 13:
		fmt.Println("child")
	//  - "teenager" when age is 13 through 17
	case age < 18:
		fmt.Println("teenager")
	//  - "adult" when age is 18+
	default:
		fmt.Println("adult")
	}
}
