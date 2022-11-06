package main
import "fmt"

// Update getLikesAndShares to return two ints
func getLikesAndShares(postId int) (int, int){
  var likesForPost, sharesForPost int
  switch postId {
  case 1:
    likesForPost = 5
    sharesForPost = 7
  case 2:
    likesForPost = 3
    sharesForPost = 11
  case 3:
    likesForPost = 22
    sharesForPost = 1
  case 4:
    likesForPost = 7
    sharesForPost = 9
  }
  fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
  // Add in a return for likesForPost and sharesForPost
  return likesForPost, sharesForPost
}

func main() {
  var likes, shares int
  // Update this line so the results of the function
  // get stored in "likes" and "shares"
  likes, shares = getLikesAndShares(4)
  
  if likes > 5 {
    fmt.Println("Woohoo! We got some likes.")
  }
  if shares > 10 {
    fmt.Println("We went viral!")
  }
}
