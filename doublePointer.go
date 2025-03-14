package main

import "sort"

//15
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			val := nums[i] + nums[left] + nums[right]
			if val == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if val < 0 {
				left++
			} else {
				right--
			}

		}
	}
	return ans
}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	//剪枝是必须的，因为要避免重复结果
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		//剪枝1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			//对j的剪枝，注意与i略有不同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				total := nums[i] + nums[j] + nums[right] + nums[left]
				if total == target {
					ans = append(ans, []int{nums[i], nums[j], nums[right], nums[left]})
					//剪枝2
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if total < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}

//27
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++

		}
		fast++
	}
	return slow
}
