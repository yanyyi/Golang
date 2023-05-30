package main

import (
	"context"
	"errors"
)

//---------------业务逻辑--------------------//

// 1.1 业务逻辑抽象为接口
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 1.2 实现接口
type addService struct{}

func (s addService) Sum(_ context.Context, a, b int) (int, error) {
	return a + b, nil
}

func (s addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", errors.New("至少一个字符串不为空")
	}
	return a + b, nil
}

// NewService addService的构造函数
func NewService() AddService {
	return &addService{
		// db:db
	}
}
