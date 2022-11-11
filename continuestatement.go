package main 

import "fmt"

func main(){
  var a  int = 40

  for a < 100 {
    if a == 50 {
      a = a + 1;
      continue;
    }
    fmt.Printf("The value of a; %d\n", a);
    a++;
  }
}
