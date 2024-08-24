package handlers

import (
	"github.com/gambitier/gocomm/appContext"
	"github.com/gambitier/gocomm/modules/users/proto"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
	AppContext *appContext.AppContext
}

func NewUserServiceServer(appContext *appContext.AppContext) *UserServiceServer {
	return &UserServiceServer{
		AppContext: appContext,
	}
}
