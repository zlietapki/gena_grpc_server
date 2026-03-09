package grpc_handler

import (
	"github.com/zlietapki/boilerplate/internal/domain"
	"github.com/zlietapki/microboiler_api_contracts/pkg/pb/v1"
)

type Handler struct {
	pb.UnimplementedExampleServer
	uc domain.IUsecase
}

func New(uc domain.IUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
