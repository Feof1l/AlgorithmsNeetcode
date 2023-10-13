package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))

}
func containsDuplicate(nums []int) bool {
	m := make(map[string]int)
	for _, v := range nums {
		if _, ok := m[strconv.Itoa(v)]; ok {
			return true
		} else {
			m[strconv.Itoa(v)] = 1
		}
	}
	return false
}
