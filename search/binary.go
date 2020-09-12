package search

//BinarySearch: 二分查询
func BinarySearch(arr []int, find int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) >> 1
		if find < arr[mid] {
			r = mid
		} else if find > arr[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
