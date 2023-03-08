package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name    string
	StuNum  [10]int
	CardNum [6]int
}

type Grade struct {
	Chiness int
	Math    int
	English int
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		reader.Reset(os.Stdin)

		gradeMap := make(map[Student]Grade)
		var s1 Student
		fmt.Println("请输入姓名:")
		fmt.Scan(&s1.Name)

		fmt.Println("请输入学号:")
		var input string
		fmt.Scan(&input)
		for i := 0; i < 10; i++ {
			num, _ := strconv.ParseInt(string(input[i]), 10, 64)
			s1.StuNum[i] = int(num)
		}

		fmt.Println("请输入卡号:")
		fmt.Scan(&input)
		for i := 0; i < 6; i++ {
			num, _ := strconv.ParseInt(string(input[i]), 10, 64)
			s1.CardNum[i] = int(num)
		}
		var grade Grade
		fmt.Println("请输入语文成绩:")
		fmt.Scan(&grade.Chiness)

		fmt.Println("请输入数学成绩:")
		fmt.Scan(&grade.Math)

		fmt.Println("请输入英语成绩:")
		fmt.Scan(&grade.English)
		gradeMap[s1] = grade

		fmt.Println(s1.Name, "的语文成绩是:", grade.Chiness)
		fmt.Println(s1.Name, "的数学成绩是:", grade.Math)
		fmt.Println(s1.Name, "的英语成绩是:", grade.English)
	}

}
