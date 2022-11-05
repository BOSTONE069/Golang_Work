package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn := true
  eludedGuards := rand.Intn(100)
  openedVault := rand.Intn(100)
  leftSafely := rand.Intn(5)


  if isHeistOn {
    if eludedGuards >= 50 {
      fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
      } else {
        fmt.Println("Plan a better disguise next time?")
        }
        if isHeistOn == true && openedVault >= 70{
          fmt.Println("Grab and Go!")
          } else if isHeistOn :=false; isHeistOn{
            fmt.Println("The Vault Cant be Opened")
            }
          switch leftSafely{
            case 1:
             isHeistOn = false
            fmt.Println("The heist is unsuccessful")
            case 2: 
             isHeistOn = false 
             fmt.Print("Turns out vault doors don't open from the inside...")
             default:
             fmt.Println("Start the gateaway car!")
          }
  }
    if isHeistOn {
      amtStolen := 10000 + rand.Intn(1000000)
      fmt.Println("The amount tha was stolen is:", amtStolen)
    }
  fmt.Println("isHeistOn is currently:", isHeistOn)

}
