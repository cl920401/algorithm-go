package search

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。
func Upstairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	step1 := 1
	step2 := 2
	for i := 3; i <= n; i++ {
		step2, step1 = step1+step2, step2
	}
	return step2
}

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func Combine(k int, n int) [][]int {
	if k > n || k <= 0 || n <= 0 {
		return [][]int{}
	}
	if k == 1 {
		arr := make([][]int, n)
		for i := range arr {
			arr[i] = []int{i + 1}
		}
		return arr
	}
	arr1 := Combine(k-1, n-1)
	arr2 := Combine(k, n-1)

	for i := range arr1 {
		arr1[i] = append(arr1[i], n)
	}
	return append(arr1, arr2...)
}
