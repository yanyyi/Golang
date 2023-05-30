package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//---------------表示对外提供的一个方法--------------------//

/*Endpoint is the fundamental building block of servers and clients.
It represents a single RPC method.
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)*/

// 1.3请求和响应
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// 2. endpoint
// 借助 适配器 将 方法 -> endpoint
func makeSumEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SumRequest)
		v, err := srv.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ConcatRequest)
		v, err := srv.Concat(ctx, req.A, req.B)
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}
