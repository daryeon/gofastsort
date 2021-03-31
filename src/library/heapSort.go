package library

/**
  堆排序
  最坏最好平均 O(NlogN)
  每个数组都能转化为堆 堆其实就是一种完全二叉树，分为大根堆和小跟堆
  1>先把数组转化为大根堆
  2>一次缩小堆长度，将大根堆顶弹出，堆尾换到堆顶进行微调 直至堆长度为0
  3 4 5 4 7 9 3
        9
      5   7
    3  4 4 3
*/
func HeapSort(arr []int) {
	//先把数组转化为大根堆
	for i := 0; i < len(arr); i++ {
		//（i-1)/2是父节点 如果比这个大 就交换
		for i > 0 && arr[i] > arr[(i-1)/2] {
			arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
			i = (i - 1) / 2
		}
	}
	//此时arr已经转化为大根堆 开始首尾交换 [9 8 7 8 7 5 4 7 7 6 6 3 5 4 3 2 5 2 6 1]
	heapSize := len(arr)
	for heapSize > 0 {
		//首尾交换
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		//堆长度-1
		heapSize--
		//调整堆为大根堆 自上而下 子节点是 i\*2+1
		//还有子节点
		curent := 0
		for curent*2+1 < heapSize {
			large := 0
			//两个子节点比较出最大的 避免越界
			if curent*2+2 < heapSize && arr[curent*2+1] < arr[curent*2+2] {
				large = curent*2 + 2
			} else {
				large = curent*2 + 1
			}
			//比较父子节点，让小数下沉
			if arr[large] < arr[curent] {
				large = curent
			}
			//看看比较出的large是不是当前节点，如果是的话说明已经最大就跳出，不是就交换位置继续下沉
			if curent != large {
				arr[curent], arr[large] = arr[large], arr[curent]
				curent = large
			} else {
				//如果父节点已经最大了就退出
				break
			}
		}
	}
}
