package main


import "fmt"



func add(number1, number2 int) int {
  return number1 + number2
}

func multiply(number1, number2 int) int {
  return number1 + number2
}

func subtract(number1, number2 int) int {
  return number1 - number2
}

func divide(number1, number2 int) float{
  if number1 == 0 || number2 == 0 {
    panic("You cannot divide by 0")
  }

  return number1 / number2
  

func main() {
  
  var action string
  var x, y, answer int

  fmt.Println("Welcome to this Golang calculator")
  fmt.Println("Please enter the first value")
  fmt.Scan(&x)
  fmt.Println("Please enter the second value")
  fmt.Scan(&y)
  fmt.Println("Please enter the action to be performed")
  fmt.Println(&action)

  if action == "add" || action == "addition" {
    add(x, y)
  } else if action == "sub" || action == "subtration" {
    subtract(x, y)
  } else if action == "mul" || action == "multiply" {
    multiply(x, y)
  } else if action == "div" || action == "divide" {
    divide(x, y)
  }

}

