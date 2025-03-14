package main

import "fmt"

func binSearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	//注意这里必须是<=不能是<
	for l <= r {
		//相比(l+r)>>1的写法能防止整型溢出
		mid := l + ((r - l) >> 1)
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

//81含重复元素的旋转数组中进行二分查找
//81相比33只多了一个清理的过程
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] && nums[mid] == nums[r] {
			l++
			r--
			//将两侧的数清理之后要重新判断l与r的关系，也要重新算mid，也就是要重新进for循环。所以下面是else if
		} else if nums[mid] > nums[r] {
			//重点就在于这个<=
			//target的范围是[l,mid)
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {

			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
