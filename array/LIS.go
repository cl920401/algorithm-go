package array

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
