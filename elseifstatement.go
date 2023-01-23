package main

import "fmt"

func main() {
	amountStolen := 64650

	// Checking if the amount stolen is greater than 1 million, if it is, it prints "We've hit the
	// jackpot!" If it is not, it checks if the amount stolen is greater than or equal to 5000, if it is,
	// it prints "Think of all the candy we can buy". If it is not, it prints "Why did we even do this?"
	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
  } else if amountStolen >= 5000{
    fmt.Println("Think of all the candy we can buy")
  } else {
		fmt.Println("Why did we even do this?")
	}

	if num := 3; num < 10{
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}
}
