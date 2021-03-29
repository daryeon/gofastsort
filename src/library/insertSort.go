package library

/**
插入排序
最坏平均 O(N^2)  最好 O(N)
像斗地主那样， 左边部分都是有序的， 右边拿一个按顺序插入左边
*/
func CharuSort(arr []int) {
	//从1开始，默认有1个有序
	for i := 1; i < len(arr); i++ {
		//比较下一个位置，往前面有序数组中插入
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}