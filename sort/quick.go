package sort

//QuickSort: 快速排序
func QuickSort(arr []int) {
	head, tail := 0, len(arr)-1
	if head >= tail {
		return
	}
	base := arr[0]
	for i := 1; i <= tail; {
		if base >= arr[i] {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		} else if base < arr[i] {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
}
