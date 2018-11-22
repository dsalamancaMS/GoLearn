package main

import (
	"fmt"
)

func main() {

	mySlices()
	printMySlice()
	printSliceValues()
}

func mySlices() {
	//make creates a slice <type>, <lenght>, <cap> lenght is the lenght of the slice, cap, is the max of the array underneath the slice
	//printf is formatting, %d acomodates intigers
	mycourses := make([]string, 5, 10)
	fmt.Printf("Lenght is: %d \nCapacity is %d \n", len(mycourses), cap(mycourses))
}

func printSliceValues() {
	mySlice := []int{2, 4, 5, 7, 1, 3}

	for _, i := range mySlice {
		fmt.Println("Content on array: ", i)
	}
}

func printMySlice() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice)
}
