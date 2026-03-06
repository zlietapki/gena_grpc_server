package grpc_handler

import (
	"context"

	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
)

type Handler struct {
	pb.UnimplementedExampleServer
	uc iUsecase
}

type iUsecase interface {
	Hello(ctx context.Context, name string) (string, error)
}

func NewHandler(uc iUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
