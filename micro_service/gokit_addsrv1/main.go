package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

var (
	ErrEmptyString = errors.New("两个参数都是空字符串")
)

// 1.1.业务逻辑抽象为接口

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 1.2 实现接口

type addService struct{}

func (svc *addService) Sum(ctx context.Context, a, b int) (int, error) {
	return a + b, nil
}

func (svc *addService) Concat(ctx context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}
	return a + b, nil
}

// 1.3 请求和响应
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

/*
Go kit通过一个称为endpoint的抽象提供给了许多功能。`Endpoint`的定义如下:

	type Endpoint func(ctx context.Context,request interface{})(response interface{}, err error)

它表示单个RPC。也就是说,我们的服务接口中只有一个方法。我们将编写简单的适配器来将服务的每个方法转换为一个端点。每个适配器将诶时候一个AddService，并返回与其中一个方法对应的端点
*/
func makeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		v, err := svc.Concat(ctx, req.A, req.B) // 方法调用
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}

// 3. transport

// decode
// 请求来了之后根据 协议(HTTP、HTTP2)和编码(JSON、pb、thrift)去解析数据
func decodeSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeConcatRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// 编码
// 把响应数据 按协议和编码 返回
// w: 代表响应的网络句柄
// response: 业务层返回的响应数据
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := &addService{}
	sumHandler := httptransport.NewServer(
		makeSumEndpoint(svc),
		decodeSumRequest,
		encodeResponse,
	)
	concatHandler := httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeConcatRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.Handle("/concat", concatHandler)
	http.ListenAndServe(":8080", nil)
}
