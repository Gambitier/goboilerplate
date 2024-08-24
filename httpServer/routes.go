package httpserver

import (
	"github.com/gambitier/gocomm/modules/users/routes"
)

func (server *HttpServer) RegisterRoutes() {
	routes.RegisterUserRoutes(server.AppContext, server.App)
}
