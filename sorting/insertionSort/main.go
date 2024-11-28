package main

import "fmt"

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j] > arr[j+1] {
			temp := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = temp
			j--
		}
	}
	return arr
}

func main() {
	newArr := []int{3, 4, 2, 1, 9, 7}
	fmt.Println(InsertionSort(newArr))
}
