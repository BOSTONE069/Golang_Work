package main

import "fmt"

type student struct{
  name string
  reg_no string
  course string
  year int
}

func main() {

  fmt.Println("Structs in Golang")

  // Declaring two variables of type student.
  var student1 student
  var student2 student

 // Assigning values to the fields of the struct.
  student1.name = "Bostone Ochieng"
  student1.reg_no = "AIIQ/01986/2015"
  student1.course = "Master of Information and Knowledge Management"
  student1.year = 2015

 // Assigning values to the fields of the struct.
  student2.name = "Kevin Odhiambo"
  student2.reg_no = "AIIQ/01407/2020"
  student2.course = "Bachelor of Science Information Science"
  student2.year = 2020

 // Printing the values of the fields of the struct.
  fmt.Printf("The name of the first student is %s\n", student1.name)
  fmt.Printf("The registration number is %s\n", student1.reg_no)
  fmt.Printf("The course of the student is %s\n", student1.course)
  fmt.Printf("The year of admission is %d\n", student1.year)

  fmt.Println("")

 // Printing the values of the fields of the struct.
  fmt.Printf("The second name of the student is %s\n", student2.name)
  fmt.Printf("The registration number of the second student is %s\n", student2.reg_no)
  fmt.Printf("The course done by the student is %s\n", student2.course)
  fmt.Printf("The year of admission is: %d\n", student2.year)

  // A short hand way of declaring a struct.
  boston := student{"BOSTONE", "AIIQ/05895/2020", "MSC IKM", 2022}

  fmt.Println(boston)
  // Printing the struct in a more readable format.
  fmt.Printf("The students details are as follows: %+v\n", boston)
}
