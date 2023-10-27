package main

import "fmt"

func main() {
	matrix := [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
func searchMatrix(matrix [][]int, target int) bool {
	// проходим сначала по строкам, ищем ту , в которой будет наш искомы элемент
	// смотри на крайние элементы в строках и переходим на нужную
	// если значение лежит между крайними элементами строки, то выходим из цикла - т.к. нашли искомую строку
	Rows := len(matrix)
	Col := len(matrix[0])
	lowRow := 0
	highRow := Rows - 1
	flag := true
	var midRow int
	for lowRow <= highRow && flag {
		midRow = lowRow + (highRow-lowRow)/2
		if target > matrix[midRow][Col-1] {
			lowRow = midRow + 1
		} else if target < matrix[midRow][0] {
			highRow = midRow - 1
		} else {
			flag = false
		}
	}
	// если цикл отработал до конца, значит искомого значния вообще нет в матрице
	if flag {
		return false
	}
	// проходим непосредственно по строке - массиву и ищем элемент как обычный бинарный поиск
	lowCol := 0
	highCol := len(matrix[0])
	for lowCol <= highCol {
		midCol := lowCol + (highCol-lowCol)/2
		if target > matrix[midRow][midCol] {
			lowCol = midCol + 1
		} else if target < matrix[midRow][midCol] {
			highCol = midCol - 1
		} else {
			return true
		}

	}

	return false
}
