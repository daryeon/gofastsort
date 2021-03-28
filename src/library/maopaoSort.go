package library

/**
冒泡 和 选择排序
复杂度最高的排序 基本不用 只做演示和教学
复杂度为 选择排序最差最好平均都是O(N^2)  冒泡最好可能是O(N)
冒泡  22比较
选择  1和后面的全部比较
*/
func MaopaoSort(arr []int) {
	//不解释了， 两两比较
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
