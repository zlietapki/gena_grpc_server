// start name:top
package main

// start name:imports type:merge
import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zlietapki/boilerplate/internal/config"
	"github.com/zlietapki/boilerplate/internal/grpc_handler"
	"github.com/zlietapki/boilerplate/internal/repository"
	"github.com/zlietapki/boilerplate/internal/usecase"
	"github.com/zlietapki/boilerplate/pkg/grpcserver"
	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
	"google.golang.org/grpc"
)

// start name:main
func main() {
	cfg := config.New()

	repo := repository.New()
	uc := usecase.New(repo)

	//start name:handler type:add
	grpcHandler := grpc_handler.New(uc)

	grpcServer := grpcserver.New(cfg.GRPCListen, func(s *grpc.Server) {
		pb.RegisterExampleServer(s, grpcHandler)
	})

	grpcServerErrCh := grpcServer.Start()
	slog.Info("gRPC listening", "addr", cfg.GRPCListen)
	defer grpcServer.Stop()

	//start name:signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:

	//start name:signals_handler type:add
	case err := <-grpcServerErrCh:
		panic("gRPC server:" + err.Error())
		// start name:bottom
	}
}
