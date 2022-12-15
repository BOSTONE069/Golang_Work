package main

import "fmt"

func main() {
	fmt.Println("Usage of Maps in Golang")

	// Creating a map of string to string.
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// Creating a map of string to string.
	fmt.Println("List of all the languages", languages)
	fmt.Println("JS is the short for ", languages["JS"])

	// Deleting the key value pair of JS from the map.
	delete(languages, "JS")
	fmt.Println("List of all languages: ", languages)
}