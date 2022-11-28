package main

import "fmt"

func main() {
  fmt.Println("Welcome to learnih arreays")
  // Creating an array of 10 integers and then assigning values to each element of the array.

  var n [10] int
  var i,j int

  for i = 0; i < 10; i++ {
    n[i] = i + 100
  }

  /* output each array elemts value */
  for  j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }

 // Creating an array of 4 strings and then assigning values to each element of the array.
  var fruitList [4]string

  fruitList[0] = "Apple"
  fruitList[1] = "Tomato"
  fruitList[3] = "Peach"

  fmt.Println("The fruit list is: ", fruitList)
  fmt.Println("The fruit list is: ", len(fruitList))

  // Creating an array of 3 strings and then assigning values to each element of the array.
  var vegList = [3]string{"potato", "beans", "mushrooms"}
  fmt.Println("The veg list is: ", vegList)
  fmt.Println("The veg list is: ", len(vegList))


}
