package library

/**
经典快排
以数组最后一位数为基准，小于的放左边 等于的放中间 大于的放右边
递归  以上  左边的  和  右边的
不稳定 最坏 O(N^2)  最好和平均 O(NlogN)
随机快排  最坏最好和平均 O(NlogN)
不取最后一位而是随机一位做基准，其他不变 所以它结合概率  期望复杂度是 O(NlogN)
*/
func QuickSort(arr []int, L, R int) {
	//L不能小于R
	if L < R {
		//将数组分为左中右三部分，返回等于的区间
		partition := partition(arr, L, R)
		//递归左部分
		QuickSort(arr, L, partition[0]-1)
		//递归右部分
		QuickSort(arr, partition[1]+1, R)
	}
}

func partition(arr []int, L, R int) []int {
	//less下标之前的都是小于待比较数的 默认数组最小下标-1
	less := L - 1
	//more下标之后的都是大于待比较数的 默认数组最大下标+1
	more := R + 1
	//待比较数
	compare_num := arr[R] //默认为最后一位 随机快排  这个取个随机数即可
	//当前比较下标， 比较到哪个位置了
	cur := L
	for cur < more {
		//fmt.Printf("cur:%v,less:%v,more:%v", cur, less, more)
		if arr[cur] < compare_num {
			//如果当前数小于比较数，则less的后一位与当前元素交换位置并把less加1，当前cur后移一位
			arr[less+1], arr[cur] = arr[cur], arr[less+1]
			less++
			cur++
		} else if arr[cur] > compare_num {
			//如果当前数大于比较数，则more的前一位与当前元素互换位置并把more减1，此时cur不移动继续比较
			arr[more-1], arr[cur] = arr[cur], arr[more-1]
			more--
		} else {
			//相等的话直接往后面移动
			cur++
		}
		//fmt.Println(arr)
	}
	//返回 相等的区间 也就是less+1和more-1都是与比较数相等的
	return []int{less + 1, more - 1}
}
