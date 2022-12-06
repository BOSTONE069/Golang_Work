// Go program to illustrate how to sort
// the elements present in the slice
package main

import (
   "fmt"
   "sort"
	 "bytes"
)

func main() {

    fmt.Println("Welcvome to the Golang Slices")

    // Creating Slice
    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

    fmt.Println("Before sorting:")
    fmt.Println("Slice 1:", slc1)
    fmt.Println("Slice 2:", slc2)
    fmt.Printf("The data type for the slc1 is %T\n", slc1)

    //adding data to a slice slice
    slc1 = append(slc1, "Javascript", "C", "C++")
    fmt.Println(slc1)

    // Performing sort operation on the
    // slice using sort function
    sort.Strings(slc1)
    sort.Ints(slc2)

    fmt.Println("\nAfter sorting:")
    fmt.Println("Slice 1:", slc1)
    fmt.Println("Slice 2:", slc2)

		 // Creating and initializing
    // slices of bytes
    // Using shorthand declaration

    slice_1 := []byte{'G', 'E', 'E', 'K', 'S'}
    slice_2 := []byte{'G', 'E', 'e', 'K', 'S'}

    // Comparing slice
    // Using Compare function
    res := bytes.Compare(slice_1, slice_2)

    if res == 0 {
        fmt.Println("!..Slices are equal..!")
    } else {
        fmt.Println("!..Slice are not equal..!")
    }

    // How to remove a value based on index

    var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
    fmt.Println(courses)
    var index int = 2
    // Removing the value at index 2.
    courses = append(courses[:index], courses[index+1:]...)
    fmt.Println(courses)

}