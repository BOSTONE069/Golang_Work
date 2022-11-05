package main

import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pagenumber int
  var grade float32

  title = "Mr. GoTOSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pagenumber = 14
  grade = 6.5
  
  fmt.Println(title, "written by", writer, "and published by", publisher, "in the year", year, "The page number", pagenumber, "was played by", artist, "I can give it a grade of", grade)

  fmt.Println("")
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pagenumber = 160
  grade = 9.0

  fmt.Println(title, "written by", writer, "and published by", publisher, "in the year", year, "The page number", pagenumber, "was played by", artist, "I can give it a grade of", grade)
}
