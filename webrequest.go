package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

//Declaring the url as a  constant global variable
const url = "https://bostone069.github.io/portforlio/index.html"

func main() {
  fmt.Println("This is a web request using Golang")

  response, err := http.Get(url)

  //Helps in error incase the link is not working well
  if err != nil{
    panic(err)
  }

  fmt.Println(response)

  defer response.Body.Close() //This is for closing the connection 

  //This is for reading the body of the response
  databytes, err := ioutil.ReadAll(response.Body)

  if err != nil {
    panic(err)
  }
  
  //This is for printing out the contents of the file in a string format
  content := string(databytes)

  fmt.Println(content)


}
