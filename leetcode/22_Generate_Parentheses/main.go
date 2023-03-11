package main

func generateParenthesis(n int) []string {
	var result []string
	result = dfs("", n, n, result)
	return result
}

func dfs(path string, left int, right int, res []string) []string {
	if left == 0 && right == 0 {
		res = append(res, path)
		return res
	}
	if left < 0 || right < 0 || left > right {
		return res
	}
	res = dfs(path+"(", left-1, right, res)
	res = dfs(path+")", left, right-1, res)
	return res
}
