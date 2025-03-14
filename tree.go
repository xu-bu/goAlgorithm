package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var dfs func(node *TreeNode)
	items, ans := []int{}, [][]int{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		target -= node.Val
		items = append(items, node.Val)
		if node.Left == nil && node.Right == nil && target == 0 {
			//ans = append(ans, items)
			ans = append(ans, append([]int(nil), items...))
		}
		dfs(node.Left)
		dfs(node.Right)
		items = items[:len(items)-1]
	}

	dfs(root)
	return ans
}

//98
func isValidBST(root *TreeNode) bool {
	nums := []int{}
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	pre := nums[0]
	for _, v := range nums {
		if v < pre {
			return false
		}
		pre = v
	}
	return true
}

func generateTrees(n int) []*TreeNode {
	ans := []*TreeNode{}
	var build func(start, end int) []*TreeNode
	build = func(start, end int) []*TreeNode {
		ret := []*TreeNode{}
		if start >= end {
			ret = append(ret, nil)
		} else if start == end-1 {
			ret = append(ret, &TreeNode{start, nil, nil})
		} else {
			for i := start; i < end; i++ {
				for _, left := range build(start, i) {
					for _, right := range build(i+1, end) {
						root := &TreeNode{i, nil, nil}
						root.Left = left
						root.Right = right
					}
				}
				ret = append(ret, root)
			}
		}
		return ret
	}
	for i := 1; i <= n; i++ {

		for _, left := range build(1, i) {
			for _, right := range build(i+1, n+1) {
				root := &TreeNode{i, nil, nil}
				root.Left = left
				root.Right = right
				ans = append(ans, root)
			}
		}
	}
	return ans
}

//99
func recoverTree(root *TreeNode) {
	nodes := []*TreeNode{}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		nodes = append(nodes, root)
		inOrder(root.Right)
	}
	inOrder(root)
	//两个降序对则交换（第一对的pre，第二对的next）
	//一个降序对则交换这一对数本身
	index := []int{}
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			index = append(index, i)
		}
	}
	if len(index) == 1 {
		nodes[index[0]].Val, nodes[index[0]+1].Val = nodes[index[0]+1].Val, nodes[index[0]].Val
	} else {
		nodes[index[0]].Val, nodes[index[1]+1].Val = nodes[index[1]+1].Val, nodes[index[0]].Val
	}
}

//110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(getHeight(root.Left)-getHeight(root.Right)) <= 1
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}

//113
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	items := []int{}
	var dfs func(target int, root *TreeNode)
	dfs = func(target int, root *TreeNode) {
		if root == nil {
			return
		}
		items = append(items, root.Val)
		target -= root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int(nil), items...))
		}
		dfs(target, root.Left)
		dfs(target, root.Right)
		items = items[:len(items)-1]
	}
	dfs(targetSum, root)
	return ans
}

//538 bst->累加树
//右中左，反中序遍历bst，会得到递减序列
func convertBST(root *TreeNode) *TreeNode {
	var pre int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Right)
		node.Val += pre
		pre = node.Val
		inOrder(node.Left)
	}
	inOrder(root)
	return root
}

