package grpcserver

import (
	"github.com/gambitier/gocomm/modules/users/handlers"
	"github.com/gambitier/gocomm/modules/users/proto"
)

func (server *GrpcServer) RegisterServices() {
	proto.RegisterUserServiceServer(server.App, handlers.NewUserServiceServer(server.AppContext))
}
