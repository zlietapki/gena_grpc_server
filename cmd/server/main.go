// start name:top
package main

// start name:imports type:merge
import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
	"google.golang.org/grpc"

	"github.com/zlietapki/gena/internal/config"
	"github.com/zlietapki/gena/internal/grpc_handler"
	"github.com/zlietapki/gena/internal/usecase"
	"github.com/zlietapki/gena/pkg/grpcserver"
)

// start name:main
func main() {
	cfg := config.New()

	// start name:usecase_deps type:add

	// start name:new_usecase
	uc := usecase.New(usecase.Depends{
		// start name:usecase_deps_objs type:add
		// start name:post_usecase_deps_objs
	})

	//start name:after_usecase type:add
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

	//start name:signals_select type:add
	case err := <-grpcServerErrCh:
		panic("gRPC server:" + err.Error())
		//start name:post_signals_select
	}
	// start name:bottom
}
