package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4}, 9))
}

// асимптотическая сложность О(log(n)), n - размерсвходного массива
func binarySearch(nums []int, target int) int { // подаем на вход отсортированный массив элементов и число, необходиоме найти
	// возвращаем позицию числа

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		// аналог для (low + high)/2 , но так не будет переполнения стека при большой длине массива( когда мы складываем большие числа может переполниться стек)
		// а так мы складываем разницу между левым и правым концом,деленную на 2, с левым концом.То есть просто левый конец + половина, то что нам и нужно, получим середину длины
		find := nums[mid]
		if find < target {
			low = mid + 1
		} else if find > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
