package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))

}
func twoSum2(nums []int, target int) []int { // мое решение 7мс, 4,3 Мб
	m := make(map[int]int)
	array := []int{}
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			if _, ok := m[target-num]; ok {
				array = append(array, i)
				array = append(array, m[target-num])
				return array
			}
			m[num] = i
		} else if num*2 == target {
			array = append(array, i)
			array = append(array, m[target-num])
		}

	}
	return array
}
func twoSum(nums []int, target int) []int { // самое эффективное решение 3 мс, 4,3 Мб, асимптотика О(n)  для врмени и памяти
	m := make(map[int]int)
	array := []int{}
	for i, num := range nums {
		diff := target - num
		if _, ok := m[diff]; ok {
			return []int{m[diff], i}
		}
		m[num] = i
	}
	return array
}
