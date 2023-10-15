package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	sl := make([][]int, len(nums)+1)
	// записали в мапу значения как key это сам элемент,
	//а значение - это сколько раз он встречается в исходном массиве
	for _, n := range nums {
		m[n] += 1
	}
	fmt.Println(m)
	//формируем массив  в матрице,добавляя элементы,где индекс - количество повторений числа, а занчение - само число
	for n, c := range m {
		sl[c] = append(sl[c], n)
	}
	fmt.Println(sl)
	res := []int{}
	// проходимся по матрице с конца(т.к. в конце элементы, которые чаще всего встречаются) и добавляем в
	// результирующий массив, пока не станет ранво k
	for i := len(sl) - 1; i > 0; i-- {
		for _, n := range sl[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}

	}
	return res
}
