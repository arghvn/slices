package main

import (
	"fmt"
)

func main() {

	mySlice := []string{"Apples", "Oranges", "Bananas"}

	fmt.Printf("Initial slice values: %v\n", mySlice)

	myFunction(mySlice)

	fmt.Printf("Final slice values: %v\n", mySlice)

}

func myFunction(fruits []string) {

	// Change Oranges to Strawberries

	fruits[1] = "Strawberries"

	fmt.Printf("Slice values in myFunction(): %v\n", fruits)

}

// Output :
// Initial slice values: [Apples Oranges Bananas]
// Slice values in myFunction(): [Apples Strawberries Bananas]
// Final slice values: [Apples Strawberries Banana

//more details :
//In fact, if we defined the array with a value but did not specify the number of cells,
// it means that we used a slice and not an array, and the method of building without a value is done with the help
//of the make internal function.
