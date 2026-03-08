package main

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func smallestBalancedIndex(nums []int) int {
    n:=len(nums)
	if (n==1){
		return -1
	}else if(n==2){
		if nums[1]==0{
			return 0
		}else if nums[0]==1{
			return 1
		}		
        return -1
	}
	left,right:=nums[0],1
	for i:=2;i<n;i++{
		right*=nums[i]
	}
	for i:=0;i<n;i++{
		if (i==15){
			fmt.Println(left,right)
		}
		if (i==0){
			if right==0{
				return 0
			}
		}else if(i==n-1){
			if left==1{
				return n-1
			}
		}else if left==right{
			return i
		}else{
			right/=nums[i+1]
			left+=nums[i]
		}
	}
	return -1
}

func main(){
	fmt.Println(smallestBalancedIndex([]int{999,818,984,995,841,822,984,978,960,997,896,926,759,961,1000,562,1,1,1,87,4,1,40}))
}