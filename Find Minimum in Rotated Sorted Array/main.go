package main

func main() {

}
func findMin(nums []int) int {
	min := nums[0]
	left := 0
	right := len(nums) - 1
	var mid int
	// ставим обычные границы для бинарного поиска
	// если так совпало,что массив остортирован,сравниваем с текущим минимумом
	// и выходим из цикла
	for left <= right {
		if nums[left] <= nums[right] {
			if nums[left] < min {
				return nums[left]
			}
		}
		// далее считаем середину  как обычно
		// если вдруг центральный элемент оказался меньше текущего минимума
		// делаем замену
		mid = left + (right-left)/2
		if nums[mid] < min {
			min = nums[mid]
		}
		// если центральный элеемент больше левого,значит у нас часть от левого
		// до центрального отсортирована, и надо искать справа,поэтому юудем двигаться вправо
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1

		}
	}
	return min
}
