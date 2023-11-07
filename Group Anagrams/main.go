package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}
func groupAnagrams(strs []string) [][]string { // 16 мс, память 7.8 МБ
	m := make(map[[26]int][]string) // создаем мапу, ключи это массивы длинной 26,т.к. в алфавите 26 символов,значения - слова
	for _, value := range strs {
		// будем проходить по словам в массиве
		// и для каждого слова подсчитвать количество букв каждого вида и записывать в массив
		array := [26]int{}
		for i := 0; i < len(value); i++ {
			array[rune(value[i])-rune('a')] += 1 //
		}
		// затем добавляем в мапу такое слова
		// далее, если будет находится анаграмма для слова, т.е. у слова будет массив с таким же набором
		// будем к ключу массив адобавлять значения слова(у нас значение - массив строк)
		m[array] = append(m[array], value)

	}
	//просто записываем значения мапы в двумерный массив
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
