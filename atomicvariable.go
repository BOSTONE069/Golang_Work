package main
import (
   "sync"
   "time"
   "math/rand"
   "fmt"
   "sync/atomic"
)
var wait sync.WaitGroup
var count int64
// The function increment() takes a string as an argument and increments the global variable count by 1
func  increment(s string)  {
   for i :=0;i<10;i++ {
      time.Sleep(time.Duration((rand.Intn(3)))*time.Millisecond)
      atomic.AddInt64(&count,1)
      fmt.Println(s,i,"Count ->",count)
   }
   wait.Done()
}
// The function increment() is called twice, once with the argument "foo: " and once with the argument
// "bar: ".
//
// The function increment() is called as a goroutine, so the two calls to increment() happen
// concurrently.
//
// The function increment() increments the variable count, and prints the value of count and the
// argument to the function.
//
// The function main() waits for the two goroutines to finish, and then prints the value of count.
//
// The output of the program is:
//
// foo: 1
// bar: 2
// foo: 3
// bar: 4
// foo: 5
// bar: 6
// foo: 7
// bar: 8
// foo: 9
// bar: 10
// last count value 10
func main(){
   wait.Add(2)
   go increment("foo: ")
   go increment("bar: ")
   wait.Wait()
   fmt.Println("last count value " ,count)
}
