package main

import "fmt"

func largestAltitude(gain []int) int {
	maxAltitude := 0
	currentAltitude := 0

	for _, g := range gain {
		currentAltitude += g
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}

func test(gain []int, expected int) {
	output := largestAltitude(gain)
	fmt.Printf("largestAltitude(%v) = %v; expected %v\n", gain, output, expected)
	if output != expected {
		fmt.Printf("Test failed")
	} else {
		fmt.Println("Test passed")
	}
}

func main() {
	test([]int{-5, 1, 5, 0, -7}, 1)
	test([]int{-4, -3, -2, -1, 4, 3, 2}, 0)
}
