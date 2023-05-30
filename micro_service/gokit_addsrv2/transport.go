package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"gokit_addsrv2/proto"
	"golang.org/x/time/rate"
	"net/http"
)

//Transport(传输层)：定义了请求和响应的格式，和如何将请求和响应编码和解码。
// HTTP JSON
// decode
// 请求来了之后根据 协议(HTTP、HTTP2)和编码(JSON、pb、thrift) 去解析数据

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

func NewHTTPServer(svc AddService, logger log.Logger) http.Handler {
	// HTTP JSON服务
	sum := makeSumEndpoint(svc)
	// go-kit/log
	//log.With(logger, "method", "sum") 派生子logger的效果
	sum = loggingMidleware(log.With(logger, "method", "sum"))(sum)
	//使用限流中间件
	sum = rateMiddleware(rate.NewLimiter(1, 1))(sum)
	sumHandler := httptransport.NewServer(
		sum, //日志中间件包一层的sum endpoint
		decodeSumRequest,
		encodeResponse,
	)

	concatHandler := httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeConcatRequest,
		encodeResponse,
	)

	// github.com/gorilla/mux
	// r := mux.NewRouter()
	// r.Handle("/sum", sumHandler).Methods("POST")
	// r.Handle("/concat", concatHandler).Methods("POST")

	//  gin
	r := gin.Default()
	r.POST("/sum", gin.WrapH(sumHandler))
	r.POST("/concat", gin.WrapH(concatHandler))

	return r
}

// gRPC
type grpcServer struct {
	proto.UnimplementedAddServer

	sum    grpctransport.Handler
	concat grpctransport.Handler
}

// NewGRPCServer 构造函数
func NewGRPCServer(svc AddService) proto.AddServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			makeSumEndpoint(svc), // endpoint
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			makeConcatEndpoint(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

func (s grpcServer) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.SumResponse), nil
}

func (s grpcServer) Concat(ctx context.Context, req *proto.ConcatRequest) (*proto.ConcatResponse, error) {
	_, resp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.ConcatResponse), nil
}

// decodeGRPCSumRequest 将Sum方法的gRPC请求参数转为内部的SumRequest
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.SumRequest)
	return SumRequest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCConcatRequest 将Concat方法的gRPC请求参数转为内部的ConcatRequest
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.ConcatRequest)
	return ConcatRequest{A: req.A, B: req.B}, nil
}

// encodeGRPCSumResponse 封装Sum的gRPC响应
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(SumResponse)
	return &proto.SumResponse{V: int64(resp.V), Err: resp.Err}, nil
}

// encodeGRPCConcatResponse 封装Concat的gRPC响应
func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ConcatResponse)
	return &proto.ConcatResponse{V: resp.V, Err: resp.Err}, nil
}
