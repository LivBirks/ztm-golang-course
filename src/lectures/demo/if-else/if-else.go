package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 scored higher")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 scored higher")
	} else {
		fmt.Println("Quiz 1 & Quiz 2 scored the same")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grades")
	} else if average(quiz1, quiz2, quiz3) < 7 {
		fmt.Println("Grades are poor")
	}
}

