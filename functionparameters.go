package main
import "fmt"

// Update marsYear so that it takes earthYears
// As a parameter
func computeMarsYears(earthYears int) int {
  // Remove earthYears definition within marsYear
  
  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears
}

func main() {
  myAge := 25
  
  // Call `marsYear` with `myAge`
  myMartianAge := computeMarsYears(myAge)
  fmt.Println(myMartianAge)
}
