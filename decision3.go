package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	i := 2 // Удаляем элемент под номером i, порядок элементов не изменяется, работает медленнее поскольку создает новый срез

	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)
}
