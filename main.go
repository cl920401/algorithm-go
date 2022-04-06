package main

import (
	"algorithm-go/search"
	"algorithm-go/sort"
	"fmt"
)

// 测试代码
func main() {
	arr := []int{2868, 5485, 1356, 1306, 6017, 8941, 7535, 4941, 6331, 6181}
	fmt.Printf("arr := %v", arr)
	fmt.Print("\r\n")
	searchNum := 6017
	fmt.Printf("search.BinarySearch(arr, %d):", searchNum)
	fmt.Print(search.BinarySearch(arr, 6017))
	fmt.Print("\r\n")

	fmt.Printf("sort.QuickSort(arr):")
	sort.QuickSort(arr)
	fmt.Print(arr)
}
