package grpcserver

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/gambitier/gocomm/appContext"
	"github.com/gambitier/gocomm/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	App        *grpc.Server
	AppContext *appContext.AppContext
}

func NewGrpcServer(appContext *appContext.AppContext) *GrpcServer {
	return &GrpcServer{
		App: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				recoveryInterceptor(),
				middlewares.ValidationInterceptor(),
			),
		),
		AppContext: appContext,
	}
}

func recoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n%s", r, debug.Stack())
				err = status.Errorf(codes.Internal, "Internal server error")
			}
		}()
		return handler(ctx, req)
	}
}
