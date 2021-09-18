package NowCoder

/*
给定一个数组，请你编写一个函数，返回该数组排序后的形式。
*/

func Sort_Bubble(arr []int) []int {
	// write code here
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 快排
func Sort_Quick(arr []int, left, right int) {
	if left > right {
		return
	}
	base := arr[left]
	i, j := left, right
	for i != j {
		// 从右往左找，找到比base小的数
		for i < j && arr[j] >= base {
			j--
		}
		// 从左往右边找，直到找到比base大的数
		for i < j && arr[i] <= base {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 将基准数放到中间的位置（基准数归位）
	arr[i], arr[left] = arr[left], arr[i]

	Sort_Quick(arr, left, i-1)
	Sort_Quick(arr, i+1, right)
}

// 选择排序
func Sort_Selection(arr []int) {
	var min int
	for i := 0; i < len(arr); i++ {
		min = i
		// 找出每一轮最小的一个数，放在第一个位置
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

