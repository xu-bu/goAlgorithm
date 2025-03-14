package main

import "fmt"

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//插入排序核心思想：把第i张牌插入到合适的位置
//从i开始，用j复制i，然后不断和j-1比，只要j-1比较大，就要把j往前推
func insert(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j-1 >= 0 && nums[j-1] > nums[j]; j-- {
			swap(nums, j, j-1)
		}
	}
}
