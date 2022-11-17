package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("Welcome to the Web Server Get Request method")
  performGetRequest()
}


func performGetRequest() {
  const myurl = "https://bostone069.github.io/portforlio/index.html"

  response, err := http.Get(myurl)

  //This for error detection in the url
  if err != nil {
    panic(err)
  }
  
  //This is for closing the connection body
  defer response.Body.Close()

  //This is for printing the content length anf the statuscode of the url
  fmt.Println("Status code: ", response.StatusCode)
  fmt.Println("The content length: ", response.ContentLength)

}
