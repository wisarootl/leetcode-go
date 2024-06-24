package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			prevEmpty := i == 0 || flowerbed[i-1] == 0
			nextEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if prevEmpty && nextEmpty {
				flowerbed[i] = 1
				count++
			}
		}
		if count >= n {
			return true
		}
	}
	return false
}

func test(flowerbed []int, n int, expected bool) {
	output := canPlaceFlowers(flowerbed, n)
	fmt.Printf("canPlaceFlowers(%v, %d) = %v; expected %v\n", flowerbed, n, output, expected)
	if output != expected {
		fmt.Printf("Test failed")
	} else {
		fmt.Println("Test passed")
	}
}

func main() {
	test([]int{1, 0, 0, 0, 1}, 1, true)
	test([]int{1, 0, 0, 0, 1}, 2, false)
}
