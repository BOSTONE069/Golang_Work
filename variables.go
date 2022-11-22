package main

import "fmt"

const LoginToken string = "hsvhghghgfhgchdhus" // public variable

func main() {
	var username string = "Bostone"
	fmt.Println(username)
	fmt.Printf("The data type for user name is : %T \n", username)
	fmt.Println("Hello World")

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The data type for the isLoggedIn is : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The data type for small value is : %T \n", smallVal)

	var smallFloat float32 = 67.789543267
	fmt.Println(smallFloat)
	fmt.Printf("The data type for small float is : %T \n", smallFloat)

	fmt.Println(LoginToken)
	fmt.Printf("The data type for login token is : %T \n", LoginToken)

}