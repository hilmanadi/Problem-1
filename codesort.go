package main

import (
	"fmt"
)

func main() {
	var array = numeric()
	var maxRow = array[0]
	var rows, cols int
	for _, array := range array {
		if array > maxRow {
			maxRow = array
		}
	}
	barchart(maxRow, array)
	fmt.Print("\n\n\n")

	fmt.Print("Sorted Array (Ascending) \n\n")
	barchart(maxRow, array)
	fmt.Print("\n\n\n")
	totalElemen := len(array) - 1
	proses := 1
	for i := 0; i <= totalElemen; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j+1] < array[j] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				fmt.Printf("Step ke %d Sorted Array Ascending \n\n", proses)
				barchart(maxRow, array)
				fmt.Print("\n\n\n")
				proses++
			} else {
				break
			}
		}
	}

	fmt.Print("Reverse Sorting \n\n")
	for rows = maxRow; rows > 0; rows-- { 
		for cols = totalElemen; cols >= 0; cols-- {
			if array[cols] >= rows {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for k := totalElemen; k >= 0; k-- {
		fmt.Print(array[k], " ")
	}

	fmt.Println()
}

func barchart(maxRow int, array []int) {
	var rows, cols int
	for rows = maxRow; rows > 0; rows-- {
		for cols = 0; cols <= len(array)-1; cols++ { 
			if array[cols] >= rows {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for _, array := range array {
		fmt.Print(array, " ")
	}
}

func numeric() []int{
	s:= []int{1,4,5,6,8,2}
	return s
}