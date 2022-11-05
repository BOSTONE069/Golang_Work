package main

import "fmt"

func main() {
	rightTime := true
	rightPlace := true

	// Edit this condition for the FIRST checkpoint
	if rightTime && rightPlace{
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}


	enoughRobbers := false
	enoughBags := true

	// Edit this condition for the SECOND checkpoint
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}
