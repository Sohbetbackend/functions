package main

import (
	"fmt"
	"log"
)

// contains function checks if the string x is contained in slice a and returns a boolean value
func contains(a []string, x string) bool {
	for _, val := range a {
		if val == x {
			return true
		}
	}
	return false
}

// getMax function finds the maximum integer from the parameters passed to the input
func getMax(nums ...int) int {
	if len(nums) < 1 {
		log.Fatalln("Empty")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(contains([]string{"zoo", "ice", "burger", "popcorn"}, "popcorn"))
	fmt.Println(getMax(10, 1000, 10000, 200, 450))
}
