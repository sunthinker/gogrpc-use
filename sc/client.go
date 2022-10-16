package sc

import (
	"context"
	"grpc-demo/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 自定义实体
type ClientEntity struct {
	ClientFd pb.UserInfoServiceClient
}

// 创建一个客户端连接
func (e *ClientEntity) InitClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	e.ClientFd = pb.NewUserInfoServiceClient(conn)

}

// 测试代码，根据自身业务实现
func (e *ClientEntity) ClientTest(Id string) {
	r, err := e.ClientFd.GetUserInfo(context.Background(), &pb.UserInfoRequest{Id: Id})
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	log.Printf("Response: %s", r.String())
}
