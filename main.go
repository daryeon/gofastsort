package main

import (
	"fmt"
	. "gofastsort/src/library"
)

func main() {
	arr1 := []int{3, 2, 4, 2, 1, 5, 6, 7, 5, 6, 7, 8, 5, 4, 3, 8, 9, 7, 6, 7}
	// 经典快排
	QuickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)

	// 冒泡排序
	arr2 := []int{3, 2, 4, 2, 1, 5, 6, 7, 5, 6, 7, 8, 5, 4, 3, 8, 9, 7, 6, 7}
	MaopaoSort(arr2)
	fmt.Println(arr2)

	// 选择排序
	arr3 := []int{3, 2, 4, 2, 1, 5, 6, 7, 5, 6, 7, 8, 5, 4, 3, 8, 9, 7, 6, 7}
	SelectSort(arr3)
	fmt.Println(arr3)

	// 插入排序
	arr4 := []int{3, 2, 4, 2, 1, 5, 6, 7, 5, 6, 7, 8, 5, 4, 3, 8, 9, 7, 6, 7}
	CharuSort(arr4)
	fmt.Println(arr4)
}
