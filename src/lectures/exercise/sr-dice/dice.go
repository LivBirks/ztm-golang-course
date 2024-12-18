//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func rollDice(numSides int) int {
    // Roll the dice and return the sum of the dice
    sumOfDice := rand.Intn(numSides) + 1
    return sumOfDice
}

func main() {
    // Ask user for number of dice, number of rolls, and number of sides
    var numDice, numRolls, numSides int
    fmt.Print("Enter number of dice: ")
    fmt.Scan(&numDice)
    fmt.Print("Enter number of rolls: ")
    fmt.Scan(&numRolls)
    fmt.Print("Enter number of sides: ")
    fmt.Scan(&numSides)

}

