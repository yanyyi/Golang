package main

import (
	"testing"
	"time"
)

// Test the decode function
func TestDecode(t *testing.T) {
	//调用被测试的函数
	post, err := decode("./post.json")
	if err != nil {
		t.Error(err)
	}
	//检查结果是否和预期的一样,如果不一样就显示一条出错信息
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got ", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World' but got ", post.Content)
	}

}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}

//func TestUnmarshal(t *testing.T)

// Test the encode function
func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
