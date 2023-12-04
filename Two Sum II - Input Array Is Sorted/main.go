package main

import "fmt"

func main() {
	numbers := []int{-1, 0}
	target := -1
	fmt.Println(twoSum(numbers, target))
}
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		//fmt.Println(numbers[left], numbers[right])
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right -= 1
		} else if sum < target {
			left += 1
		}
	}
	return []int{0, 0}
}
