
package main

//slices are a key data type in the Golang that provide a strong interface for managing arrays. 
//In fact, if you want arrays to have more features and be more flexible, you have to use slices

//Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

import "fmt"

//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
//To create an empty slice with non-zero length, use the builtin 'make'. 
//Here we make a slice of 'string's of length '3' (initially zero-valued)

func main() 

 s := make([]string, 3)

fmt.Println("emp:", s)

//We can set and get just like with arrays.


s[0] = "a"

s[1] = "b"

s[2] = "c"

fmt.Println("set:", s)

fmt.Println("get:", s[2])

//'len' returns the length of the slice as expected.

fmt.Println("len:", len(s)






//In addition to these basic operations, slices support several more that make them richer than arrays.
// One is the builtin 'append', which returns a slice containing one or more new values.
// Note that we need to accept a return value from 'append' as we may get a new slice value.


s = append(s, "d")

s = append(s, "e", "f")

fmt.Println("apd:", s)



// Slices can also be 'copy'd. Here we create an empty slice 'c 'of the same length as 's' and copy into 'c' from 's'.


c := make([]string, len(s))

 copy(c, s)

 fmt.Println("cpy:", c)

 //Slices support a "slice" operator with the syntax 'slice[low:high]'.
 // For example, this gets a slice of the elements 's[2]', 's[3]', and 's[4]'.

 
 l := s[2:5]

 fmt.Println("sl1:", l)

 // This slices up to (but excluding) 's[5]'.

 l = s[:5]

 fmt.Println("sl2:", l)

//And this slices up from (and including) 's[2]'.

 l = s[2:]

 fmt.Println("sl3:", l)
 
//We can declare and initialize a variable for slice  in a single line as well.

 t := []string{"g", "h", "i"}

 fmt.Println("dcl:", t)

 // Slices can be composed into multi-dimensional data structures. 
 //The length of the inner slices can vary, unlike with multi-dimensional arrays.

 
 twoD := make([][]int, 3)

 for i := 0; i < 3; i++ {

 innerLen := i + 1

 twoD[i] = make([]int, innerLen)

 for j := 0; j < innerLen; j++ {

 twoD[i][j] = i + j

 }

 }

 fmt.Println("2d: ", twoD)
 
}


// output :
//emp: [ ]
//set: [a b c]
//get: c
//len: 3
//apd: [a b c d e f]
//cpy: [a b c d e f]
//sl1: [c d e]
//sl2: [a b c d e]
//sl3: [c d e f]
//dcl: [g h i]
//2d: [[0] [1 2] [2 3 4]]


// more details :
//Unlike arrays, slices are not always defined with a fixed length and depend on the values poured into them.

//The make function can be used to create an empty and valueless slice.
//In the first example, we made a size 3 slice without any values.
//In slices, like arrays, you can select values for each house and call them.
//The len dimension also returns the cut length like arrays.
//Slices have more capabilities than arrays. For example, the append function can be used to add one or more values to a slice.
//To copy a slice, we can use the copy command in the definition of the new slice on the variable c. 
//First, a slice is created with the desired slice size, and then the values inside it are copied.
//Slices use syntax [to get the part we want from inside the slice. Such as variable l.
//The t variable shows how to define a slice with values in a line.
//Slices can be implemented in a multidimensional data structure, but the length of the values inside the slices can be different and not like multidimensional arrays.
//Unlike arrays, slides act as references when sent to a function as a parameter, and changes to them cause the original array to change.