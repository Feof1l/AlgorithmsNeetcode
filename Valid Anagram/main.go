package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "cat"

	t := "tac"
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram2(s, t))
	fmt.Println(isAnagramSort(s, t))

}
func isAnagram(s string, t string) bool { // 2 по скорости 9 мс , память 2.4 МБ
	m := make(map[rune]int)
	for _, char := range s {
		if _, ok := m[char]; !ok {
			m[char] = 1
		} else {
			m[char] += 1
		}

	}
	for _, char := range t {
		if _, ok := m[char]; !ok {
			return false
		} else {
			m[char] -= 1
			if m[char] < 0 {
				return false
			}
		}

	}
	for _, value := range m {
		if value > 0 {
			return false
		}
	}
	return true
}
func isAnagram2(s string, t string) bool { // самое быстрое решение 8 мс, память 2.4 МБ
	if len(s) != len(t) {
		return false
	}
	mS := make(map[rune]int)
	mT := make(map[rune]int)
	for _, char := range s {
		if _, ok := mS[char]; !ok {
			mS[char] = 1
		} else {
			mS[char] += 1
		}

	}
	for _, char := range t {
		if _, ok := mT[char]; !ok {
			mT[char] += 1
		} else {
			mT[char] += 1
		}

	}
	for key, value := range mS {
		if value != mT[key] {
			return false
		}
	}
	return true
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func isAnagramSort(s string, t string) bool { // худшее решение 22мс , память 6.6 мб, но если взять сортировку за О(1), то по идее должно быть лучше
	return SortString(s) == SortString(t)
}
