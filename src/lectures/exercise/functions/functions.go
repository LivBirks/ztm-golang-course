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

import (
	"fmt"
	"math/rand"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) string {
	return "Hi" + " " + name
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func giveMessage(message string) string {
	return message
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func randomNumber() int {
	return rand.Intn(10)
}

//* Write a function that returns any two numbers
func twoRandomNumbers() (int, int) {
	return rand.Intn(10), rand.Intn(10)

}

func main() {
	//* Add three numbers together using any combination of the existing functions.
	sumOfThreeNumbers := addThreeNumbers(1, 2, 3)
	//  * Print the result
	fmt.Println(sumOfThreeNumbers)
	//* Call every function at least once
	fmt.Println(greet("Liv"), giveMessage("Great work"), randomNumber())
	num1, num2 := twoRandomNumbers()
	print(num1, " ", num2)

}

