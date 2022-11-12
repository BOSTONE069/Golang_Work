package main

import "fmt"

type Student struct{
  name string
  reg_no string
  course string
  year int
}

func main() {
  var student1 Student
  var student2 Student

  student1.name = "Bostone Ochieng"
  student1.reg_no = "AIIQ/01986/2015"
  student1.course = "Master of Information and Knowledge Management"
  student1.year = 2015

  student2.name = "Kevin Odhiambo"
  student2.reg_no = "AIIQ/01407/2020"
  student2.course = "Bachelor of Science Information Science"
  student2.year = 2020

  studentDetails(&student1)

  fmt.Println("")

  studentDetails(&student2)
}

func studentDetails( student *Student) {
  fmt.Printf("The name of the studennt is: %s\n", student.name)
  fmt.Printf("The registration number of the student is: %s\n", student.reg_no)
  fmt.Printf("The course that is taken by the student is: %s\n", student.course)
  fmt.Printf("The year of admission of the student is: %d\n", student.year)
}
