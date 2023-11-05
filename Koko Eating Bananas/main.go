package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(piles, h))

}
func Max(piles []int) int {
	max := -1
	for i := range piles {
		if piles[i] > max {
			max = piles[i]
		}
	}
	return max
}

// создаем массив от 1 до макс элементав исходном массиве.
// бин поиск будем проводить в этом массиве
// середина - скорость поедания бананов
// для этой скорости будем высчитывать количество часов
// если оно больше заданного кол ва часов, то двигаем левую границу
// и в результат записываем меньшее из к и предыдущий результат
// если оно меньше , двигаем правую границу
func minEatingSpeed(piles []int, h int) int {

	left := 1
	right := 1000000000
	res := right
	for left <= right {
		k := left + (right-left)/2
		hours := 0
		for _, p := range piles {

			hours += int(math.Ceil((float64(p) / float64(k))))
		}
		if hours <= h {
			res = int(math.Min(float64(res), float64(k)))
			right = k - 1
		} else {
			left = k + 1
		}

	}
	return res

}
