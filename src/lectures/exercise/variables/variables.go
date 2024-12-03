//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	// using var and assuming type
	var favouriteColour = "Blue"
	fmt.Println("My favourite colour is", favouriteColour)
	// compound assignment
	birthYear, ageInYears := 1999, 25
	// fmt.Println("My birth year is", birthYear)
	// fmt.Println("My age in years is", ageInYears)
	fmt.Println("My birth year is", birthYear, "and my age in years is", ageInYears)
	// block assignment
	var (
		firstInitial = "L"
		lastInitial = "B"
	)
	fmt.Println("My initials:", firstInitial, lastInitial)
	var ageInDays int
	ageInDays = ageInYears * 365
	fmt.Println("My age in days is", ageInDays)


}

