package main

import "fmt"

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func boolValue(programming *bool) {
  *programming = false
}
func floatValue(kiswahiliMarks *float64){
  *kiswahiliMarks = 47.89
}

func main() {
	greeting := "Hello there!"
	coding := true 
  kiswahilimarks := 92.45

  brainwash(&greeting)
	boolValue(&coding)
  floatValue(&kiswahilimarks)

	fmt.Println("greeting is now:", greeting)
  fmt.Println("The application of programming in computers today is ", coding)
  fmt.Println("The updated marks for Kiswahili is", kiswahilimarks)
}
