package main

import (
	"flag"
	"fmt"

	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/config"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/server"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/conversation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewConversationServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterConversationServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
