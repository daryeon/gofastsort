package library

func SelectSort(arr []int) {
	//不解释了， 与后面的一一比较
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if tmp > arr[j] {
				tmp, arr[j] = arr[j], tmp
			}
		}
		arr[i] = tmp
	}
}