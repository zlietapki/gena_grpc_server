package grpcserver

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type addHandlerFunc func(s *grpc.Server)

type Server struct {
	listen     string
	srv        *grpc.Server
	addHandler addHandlerFunc
	opts       []grpc.ServerOption
}

func NewGRPCServer(listen string, addHandler addHandlerFunc, opts ...grpc.ServerOption) *Server {
	defaultOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(recoveryInterceptor()),
	}
	defaultOpts = append(defaultOpts, opts...)

	return &Server{
		listen:     listen,
		addHandler: addHandler,
		opts:       defaultOpts,
	}
}

func (s *Server) Start() chan error {
	s.srv = grpc.NewServer(s.opts...)
	reflection.Register(s.srv)

	hsrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.srv, hsrv)

	s.addHandler(s.srv)

	errCh := make(chan error)

	lis, err := net.Listen("tcp", s.listen)
	if err != nil {
		go func() {
			errCh <- err
		}()

		return errCh
	}

	go func() {
		errCh <- s.srv.Serve(lis)
	}()

	return errCh
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
