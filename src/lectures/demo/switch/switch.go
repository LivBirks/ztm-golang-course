package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap price")
	case p < 10:
		fmt.Println("Ok price")
	default:
		fmt.Println("Expensive price")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy ticket")
	case Business:
		fmt.Println("Business ticket")
	default:
		fmt.Println("Standard ticket")
	}

}

