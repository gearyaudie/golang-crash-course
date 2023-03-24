package main

import "fmt"

func main() {
	purchases := [5]float32{19.99, 20.99, 1.99, 2.99, 3.99}

	// this is a slice
	// mySlice := purchases[:]
	// mySlice = append(mySlice, 14.99, 15.99, 19.99)

	myOtherSlice := purchases[:3]

	myThirdSlice := purchases[2:3]

	combine := append(myOtherSlice, myThirdSlice...)

	fmt.Println(combine)
}