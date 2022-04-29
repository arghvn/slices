package main

//We saw that when using make we have to specify the number of cells or the capacity.
// To do this, however, this built-in function also has a third variable, which means maximum capacity.
// Slides have an internal function called cap that indicates the capacity of the slice!
//In the example below, a slice with an initial capacity of 4 and a maximum of 8 is made.

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 4, 8)

	fmt.Printf("Initial Length: %d\n", len(mySlice))

	fmt.Printf("Capacity: %d\n", cap(mySlice))

	fmt.Printf("Contents: %v\n", mySlice)

}

// Output :
// Initial Length: 4
// Capacity: 8
// Contents: [0 0 0 0]
