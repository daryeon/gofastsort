package library


/**
归并排序
最坏最好平均 O(NlogN)
递归
左右 两个数组 依次比较   递归出来的左右两个数组是分别有序
*/
func GuibingSort(arr []int) []int {
	//数组个数小于2直接返回
	if len(arr) < 2 {
		return arr
	}
	index := len(arr) / 2
	//归并左边的使其有序
	left := GuibingSort(arr[0:index])
	//归并右边的使其有序
	right := GuibingSort(arr[index:])
	//合并两个有序数组
	return merge(left, right)
}
func merge(left []int, right []int) []int {
	//临时数组
	tmp := make([]int, 0)
	//左右两个有序数组都从下标0开始比较
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			//如果左边小于右边，左边加入临时数组，左边下标+1继续比较
			tmp = append(tmp, left[i])
			i++
		} else {
			//如果左边大于右边，则相反
			tmp = append(tmp, right[j])
			j++
		}
	}
	//上面比较完可能有一方还有数组没放进去，则直接放进临时数组，下面两个语句只有一个会执行，可以让左边或右边剩下元素直接放到临时后面即可
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	//返回合并好的有序数组
	return tmp
}