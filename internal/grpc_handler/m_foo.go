package grpc_handler

import (
	"context"

	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Foo(ctx context.Context, req *pb.FooReq) (*pb.FooRes, error) {
	hello, err := h.uc.Hello(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "application error")
	}

	return &pb.FooRes{
		Hello: hello,
	}, nil
}
