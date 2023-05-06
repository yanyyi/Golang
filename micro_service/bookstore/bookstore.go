package main

import (
	"bookstore/proto"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// bookstore grpc服务

type server struct {
	proto.UnimplementedBookstoreServer
	bs *bookstore
}

// ListShelve 列出所有书架的RPC方法
func (s *server) ListShelves(ctx context.Context, in *empty.Empty) (*proto.ListShelvesResponse, error) {
	//调用orm操作的那些方法
	shelvesList, err := s.bs.ListShelves(ctx)
	if err == gorm.ErrEmptySlice {
		//没有数据
		return &proto.ListShelvesResponse{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "quert failed")
	}
	// 封装返回数据
	newShelvesList := make([]*proto.Shelf, 0, len(shelvesList))
	for _, s := range shelvesList {
		newShelvesList = append(newShelvesList, &proto.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	return &proto.ListShelvesResponse{Shelves: newShelvesList}, nil
}

func (s *server) CreateShelf(ctx context.Context, in *proto.CreateShelfRequest) (*proto.Shelf, error) {
	// 参数检查
	if len(in.GetShelf().GetTheme()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid theme")
	}
	data := Shelf{
		Theme: in.GetShelf().GetTheme(),
		Size:  in.GetShelf().GetSize(),
	}
	//去数据库创建
	newShelf, err := s.bs.CreateShelf(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "create failed")
	}
	return &proto.Shelf{Id: newShelf.ID, Theme: newShelf.Theme, Size: newShelf.Size}, nil
}

func (s *server) GetShelf(ctx context.Context, in *proto.GetShelfRequest) (*proto.Shelf, error) {
	// 参数检验
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	newShelf, err := s.bs.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	return &proto.Shelf{Id: newShelf.ID, Theme: newShelf.Theme, Size: newShelf.Size}, nil
}

func (s *server) DeleteShelf(ctx context.Context, in *proto.DeleteShelfRequest) (*empty.Empty, error) {
	// 参数check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	err := s.bs.DeleteShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &emptypb.Empty{}, nil
}
