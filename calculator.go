package main


import "fmt"



// The function add takes two integers as arguments and returns an integer
func add(number1, number2 int) int {
  return number1 + number2
}

// `multiply` takes two integers, `number1` and `number2`, and returns the product of the two numbers
func multiply(number1, number2 int) int {
  return number1 + number2
}

// `subtract` takes two integers, `number1` and `number2`, and returns an integer that is the result of
// subtracting `number2` from `number1`
func subtract(number1, number2 int) int {
  return number1 - number2
}

// It divides two numbers and returns the result.
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

  // Checking the value of the variable `action` and if it is equal to `add` or `addition` it will call
  // the function `add` and pass the values of `x` and `y` to it. If the value of `action` is equal to
  // `sub` or `subtration` it will call the function `subtract` and pass the values of `x` and `y` to
  // it. If the value of `action` is equal to `mul` or `multiply` it will call the function `multiply`
  // and pass the values of `x` and `y` to it. If the value of `action` is equal to `div` or `divide`
  // it will call the function `divide` and pass the values of `x` and `y` to it.
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
