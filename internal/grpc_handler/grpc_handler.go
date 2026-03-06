package grpc_handler

import (
	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
	"github.com/zlietapki/microboiler_grpc_server/internal/domain"
)

type Handler struct {
	pb.UnimplementedExampleServer
	uc domain.IUsecase
}

func NewHandler(uc domain.IUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
