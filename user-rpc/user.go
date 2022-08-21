package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"

	"zero-demo/user-rpc/internal/config"
	"zero-demo/user-rpc/internal/server"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserCenterServer(grpcServer, server.NewUserCenterServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(TestServerInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("--------------------------\n")
	fmt.Printf("req===>%+v \n", req)
	fmt.Printf("ctx===>%+v \n", ctx)
	resp, err = handler(ctx, req)
	fmt.Printf("--------------------------\n")
	fmt.Printf("info===>%+v \n", info)
	return
}
