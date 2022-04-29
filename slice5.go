package main

//Another way to add is the internal append function. One of the important things about this function is that if we build a house over capacity,
// it does not make a mistake and automatically doubles the capacity.

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 0, 8)

	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)

	fmt.Printf("Contents: %v\n", mySlice)

	fmt.Printf("Number of Items: %d\n", len(mySlice))

	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice = append(mySlice, 19)

	fmt.Printf("Contents: %v\n", mySlice)

	fmt.Printf("Number of Items: %d\n", len(mySlice))

	fmt.Printf("Capacity: %d\n", cap(mySlice))

}

//Output :
//Contents: [1 3 5 7 9 11 13 17]
//Number of Items: 8
//Capacity: 8
//Contents: [1 3 5 7 9 11 13 17 19]
//Number of Items: 9
//Capacity: 16
