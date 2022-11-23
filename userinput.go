package main

import(
  "fmt"
  "bufio"
  "os"
)



func main() {
  //This code is using the bufio and os imput to take user input
  welcome := "Welcome to dealing with user input"
  fmt.Println(welcome)

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter the name of the footaball team:")

  userInput, err := reader.ReadString('\n')

  if err != nil {
    panic(err)
  }

  fmt.Println("The football team is", userInput)


// this code below is for using the scan method to take user input
  fmt.Println("What would you like for lunch?")

  // Add your code below:
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.", food)


}
