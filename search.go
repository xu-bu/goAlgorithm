package main

//dfs矩阵
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	rows, columns, n := len(board), len(board[0]), len(word)
	//嵌套函数的写法
	var dfs func(i, j, k int) bool
	//判断从board[i][j]开始dfs能否搜索出word
	dfs = func(i int, j int, k int) bool {
		//判断i,j是否越界，board[i][j] 是否等于 word[k]
		if i < 0 || i >= rows || j < 0 || j >= columns || board[i][j] != word[k] {
			return false
		}
		if k == n-1 {
			return true
		}
		//标记board[i][j]，这样以后搜索的时候不会再搜这个位置
		board[i][j] = ' '
		res := dfs(i-1, j, k+1) || dfs(i, j-1, k+1) || dfs(i+1, j, k+1) || dfs(i, j+1, k+1)
		//上下左右都不行，说明当前位置不行，回溯，撤销board[i][j]的标记
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	status := make([]int, numCourses)
	//课程：前置课程
	dic := map[int][]int{}
	for i := 0; i < len(prerequisites); i++ {
		dic[prerequisites[i][0]] = append(dic[prerequisites[i][0]], prerequisites[i][1])
	}
	var dfs func(c int) bool
	dfs = func(c int) bool {
		if status[c] == 1 {
			return false
		}
		status[c] = 1
		for _, v := range dic[c] {
			if status[v] == 2 {
				continue
			}
			if !dfs(v) {
				return false
			}
		}
		status[c] = 2
		return true
	}
	for i := 0; i < len(prerequisites); i++ {
		if status[prerequisites[i][1]] == 2 {
			continue
		}
		if !dfs(prerequisites[i][1]) {
			return false
		}
	}
	return true
}
