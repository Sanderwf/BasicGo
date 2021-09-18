package NowCoder

import (
	"fmt"
	"testing"
)

var nums = []int{10, 6, 90, 44, 34, 67}

func TestSortSlice_Bubble(t *testing.T) {
	outs := Sort_Bubble(nums)
	fmt.Println(outs)
}

func TestSortSlice_Quick(t *testing.T) {
	Sort_Quick(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestSortSlice_Selection(t *testing.T) {
	Sort_Selection(nums)
	fmt.Println(nums)
}