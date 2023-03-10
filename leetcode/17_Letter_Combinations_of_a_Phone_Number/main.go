package main

func letterCombinations(digits string) []string {
	queue := make([]string, 0)
	numMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	for _, digit := range digits {
		//若queue长度为0,则先将digits里第一个数组对应的字符串分别压入队列
		if len(queue) == 0 {
			for i := 0; i < len(numMap[digit]); i++ {
				queue = append(queue, string(numMap[digit][i]))
			}
		} else {
			//现有队列元素一定都是等长的,获取现有队列的长度
			llen := len(queue)
			//循环开始,循环次数为现有队列长度
			for i := 0; i < llen; i++ {
				//取出现有队列的首个元素，其余元素往前移动一格
				frontStr := queue[0]
				queue = queue[1:len(queue)]
				//将取出元素与digit对应的所有字母拼接后重新加入队列末尾
				for j := 0; j < len(numMap[digit]); j++ {
					queue = append(queue, frontStr+string(numMap[digit][j]))
				}

			}
			//循环结束
		}
	}
	return queue

}
