//go:generate swag init -o ./_apidocs

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gambitier/gocomm/appContext"
	"github.com/gambitier/gocomm/config"
	grpcserver "github.com/gambitier/gocomm/grpcServer"
	httpserver "github.com/gambitier/gocomm/httpServer"
	"google.golang.org/grpc/reflection"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	appCtx, err := appContext.InitAppContext()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer appCtx.DbRepo.Close() // close once per app
	appCtx.RegisterMsgQueues()

	restApi := httpserver.NewHttpServer(appCtx)
	restApi.Configure()
	restApi.RegisterRoutes()

	// Channel to listen for errors
	errChan := make(chan error, 1)

	// Start Fiber server in a goroutine
	go func() {
		if err := restApi.App.Listen(fmt.Sprintf(":%v", appCtx.Configs.WebServerPort)); err != nil {
			errChan <- fmt.Errorf("failed to start Fiber server: %w", err)
		}
	}()

	grpcServer := grpcserver.NewGrpcServer(appCtx)
	grpcServer.RegisterServices()

	if appCtx.Configs.Environment == config.Development {
		log.Printf("Enabling gRPC reflection for %v mode", config.Development)
		reflection.Register(grpcServer.App)
	}

	// Start gRPC server in a goroutine
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%v", appCtx.Configs.GrpcServerPort))
		if err != nil {
			errChan <- fmt.Errorf("failed to listen on port %v: %w", appCtx.Configs.GrpcServerPort, err)
			return
		}
		log.Printf("gRPC server listening on port %v", appCtx.Configs.GrpcServerPort)
		if err := grpcServer.App.Serve(listener); err != nil {
			errChan <- fmt.Errorf("failed to serve gRPC server: %w", err)
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("server error: %v", err)
	case <-quit:
		log.Println("shutting down servers...")

		// Gracefully shut down the Fiber server
		if err := restApi.App.Shutdown(); err != nil {
			log.Printf("Fiber server shutdown error: %v", err)
		}

		// Gracefully shut down the gRPC server
		grpcServer.App.GracefulStop()

		log.Println("servers shut down gracefully")
	}
}
