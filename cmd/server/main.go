package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
	"github.com/zlietapki/microboiler_grpc_server/internal/config"
	"github.com/zlietapki/microboiler_grpc_server/internal/grpc_handler"
	"github.com/zlietapki/microboiler_grpc_server/internal/usecase"
	"github.com/zlietapki/microboiler_grpc_server/pkg/grpc_server"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	cfg := config.New()

	uc := usecase.New()

	grpcHandlers := grpc_handler.NewHandler(uc)

	grpcServer := grpc_server.NewGRPCServer(cfg.GRPCListen, func(s *grpc.Server) {
		pb.RegisterExampleServer(s, grpcHandlers)
	})

	grpcServerErrCh := grpcServer.Start()
	log.Printf("gRPC listen on: %s", cfg.GRPCListen)
	defer grpcServer.Stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:
	case err := <-grpcServerErrCh:
		panic("gRPC server:" + err.Error())
	}
}
