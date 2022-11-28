// Go program to illustrate how to sort
// the elements present in the slice
package main

import (
   "fmt"
   "sort"
	 "bytes"
)

func main() {

    // Creating Slice
    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

    fmt.Println("Before sorting:")
    fmt.Println("Slice 1:", slc1)
    fmt.Println("Slice 2:", slc2)

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

}