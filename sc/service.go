package sc

import (
	"context"
	"errors"
	"grpc-demo/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// 自定义实体
type ServiceEntity struct {
	// 新版proto生成的go代码中多了一个接口，由此结构体实现了，所以要引入
	pb.UnimplementedUserInfoServiceServer
	// 自身业务需要...
}

// 创建GRPC服务
func (e *ServiceEntity) InitService(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcSer := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(rpcSer, e)
	if err := rpcSer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// GRPC的GetUserInfo接口实现
// 根据自身业务实现
func (e *ServiceEntity) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	if in.GetId() == "123456789" {
		dept := "yanfa"
		response := &pb.UserInfoResponse{
			Id:       "123456789",
			Username: "demo01",
			Age:      34,
			Dept:     &dept,
			Roles:    []string{"manager", "admin"},
		}
		return response, nil
	} else if in.GetId() == "987654321" {
		dept := "yunying"
		response := &pb.UserInfoResponse{
			Id:       "987654321",
			Username: "demo02",
			Age:      33,
			Dept:     &dept,
			Roles:    []string{"root", "common"},
		}
		return response, nil
	}
	return &pb.UserInfoResponse{}, errors.New("user id not find!")
}
