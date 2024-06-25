package main

import (
	"fmt"
	"twosum/source"
)

func main() {

	// Example array and target value
	nums := []int{3, 2, 4}
	target := 6

	// Call the twoSum function and store the result
	result := source.TwoSum(nums, target)

	// Print the result
	fmt.Println("Indices of numbers that add up to the target:", result)
}
