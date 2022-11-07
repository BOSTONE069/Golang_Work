package main

import "fmt"

// Change brainwash to have a pointer parameter
func brainwash(saying *string) {
	// Dereference saying below: 
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	
	// Call your brainwash() below:
	brainwash(&greeting)
	
	fmt.Println("greeting is now:", greeting)
}
