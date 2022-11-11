package main

import "fmt"

func main() {
	var a int = 40
	
	for a < 100 {
	   fmt.Printf("The value of a; %d\n", a);
	   a++;
	   if a > 70 {
	      break;
	  }
	}
}
