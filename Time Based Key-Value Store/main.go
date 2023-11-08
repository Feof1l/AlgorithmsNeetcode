package main

func main() {

}

type TimeMap struct {
	m map[string][]ValStamp // мапа, в котрой ключ, это строковое значение из условия
	// а значением будет массив из 2 элементов [val,time]
	// для этого создадим структуру ValStamp
}

type ValStamp struct {
	Val  string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]ValStamp)}
}

// здесь будем просто проверять, есть ли значение в нашейц мапе, если нет, делаем пустой массив и записываме в него входные параметры функции
func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = make([]ValStamp, 0)
	}
	this.m[key] = append(this.m[key], ValStamp{value, timestamp})
}

// здесь сначала проверяем, есть ли значение в нашей мапе. Если есть, values становится равным значению из мапы
// потом классический бин поиск
// Проходимя по занчению времени из нашего массива values, т.к. мы знаем, что по условию, массив будет отсортирован по времени
// именно поэтому мы можем использовать бин поиск
// И т.к. нам также нужно будет получить ближайшее значение к искомому, мы будем запоминать в res значение, если
// серединный элемент меньше искомого
func (this *TimeMap) Get(key string, timestamp int) string {
	var res string
	var values []ValStamp
	if _, ok := this.m[key]; ok {
		values = this.m[key]
	}
	left := 0
	right := len(values) - 1
	for left <= right {
		mid := left + (right-left)/2
		if values[mid].Time <= timestamp {
			res = values[mid].Val
			left = mid + 1
		} else {
			right = mid - 1

		}
	}
	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
