package boot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	baziga "github.com/pibigstar/bazinga/app/grpc"
	"github.com/pibigstar/bazinga/app/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

// grpc注册
func StartGrpc() {
	if g.Cfg().Get("grpc") == nil {
		return
	}
	grpcAddr := fmt.Sprintf("%s:%d", g.Cfg().GetString("grpc.host"), g.Cfg().GetInt("grpc.port"))
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
				err = status.Errorf(codes.Internal, "%s", p)
				glog.Errorln("grpc panic", err)
				return err
			})),
		)),
	)

	pb.RegisterBazingaServiceServer(server, &baziga.BazingaService{})

	log.Println("grpc start on", grpcAddr)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
