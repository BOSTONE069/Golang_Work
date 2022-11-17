package main


import (
  "fmt"
  "net/url"
)

const myurl string = "https://bostone069.github.io/portforlio/index.html"


func main() {
  fmt.Println("Welcome to handling of URLs in Golang")
  fmt.Println(myurl)

  //parsing
  result, err := url.Parse(myurl)

  if err != nil {
    panic(err)
  }

  fmt.Println(result.Scheme)
  fmt.Println(result.Host)
  fmt.Println(result.Path)
  fmt.Println(result.Port())
  fmt.Println(result.RawQuery)
  
  //This is defining the part of a url to use
  partsOfUrl := &url.URL {
    Scheme: "https",
    Host:   "lco.dev",
    Path:   "/tutcss",
    RawPath:"user=hitesh",
  }

  //This is a variable declared to contain the part of the url
  anotherURL := partsOfUrl.String()
  fmt.Println(anotherURL)

}
