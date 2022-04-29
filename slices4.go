package main

//In text file number 4, because the maximum capacity was 8, if we try to fill 9 cells, we will encounter an error.

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 0, 8)

	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)

	fmt.Printf("Contents: %v\n", mySlice)

	fmt.Printf("Number of Items: %d\n", len(mySlice))

	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice[8] = 19

	fmt.Printf("Contents: %v\n", mySlice)

	fmt.Printf("Number of Items: %d\n", len(mySlice))

	fmt.Printf("Capacity: %d\n", cap(mySlice))

}

//Output :
//Contents: [1 3 5 7 9 11 13 17]
//Number of Items: 8
//Capacity: 8
//panic: runtime error: index out of range
