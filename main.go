package main

import (
	"algorithm-go/search"
	"algorithm-go/sort"
	"fmt"
)

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

// 最长递增子序列 LIS
func LIS(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	lis := make([]int, 0, len(arr))
	lis = append(lis, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] > lis[len(lis)-1] {
			lis = append(lis, arr[i])
		} else if arr[i] < lis[len(lis)-1] {
			l, r := 0, len(lis)-1
			for l < r {
				mid := (l + r) >> 1
				if arr[i] <= lis[mid] {
					r = mid
				} else {
					l = mid + 1
				}
			}
			lis[r], arr[i] = arr[i], lis[r]
		}
	}
	return lis
}
