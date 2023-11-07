package main

import "fmt"

func main() {
	target := -1
	nums := []int{3, 4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, target))
}
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			// если наша цель  меньше, чем левый элемент,очевидно, она может лежать только в правой части
			// и точно так же , если она больше серединного элемента, она может лежать только в правой части
			// поэтому двигаем левую границу
			if target < nums[left] || target > nums[mid] {
				left = mid + 1
				// здесь элемент лежит в диапазоне между серединным и левым элементом, поэтому двигаем правую границу
			} else {
				right = mid - 1
			}
		} else {
			// если наша цель меньше, чем серединный элемент, она должна лежать в левой части
			// и точно так же если наша цель больше правого элемента,она доолжна лежать в левой части
			// поэтому мы двигаем правую границу
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
				// здесь значит, что элемент лежит как раз в диапазоне от центрального до правого элемента
				// значит, нужно двигать левую границу
			} else {
				left = mid + 1

			}
		}

	}
	return -1
}
