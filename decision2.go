package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	i := 2 // Удаляем элемент под номером i, порядок элементов не изменяется, сложность линейная

	copy(arr[i:], arr[i+1:])

	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
