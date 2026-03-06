package usecase

import "github.com/zlietapki/microboiler_grpc_server/internal/domain"

var _ domain.IUsecase = (*UC)(nil) // compile-time check

type UC struct {
	repo domain.IRepository
}

func New(repo domain.IRepository) *UC {
	return &UC{
		repo: repo,
	}
}
