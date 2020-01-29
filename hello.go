package main

import "fmt"

func main() {
	var slice1 = make([]int, 10)
	slice2 := slice1
	slice2[1] = 12
	fmt.Print(slice1[1])
}

