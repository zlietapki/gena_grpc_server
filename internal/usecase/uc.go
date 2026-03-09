package usecase

import "github.com/zlietapki/boilerplate/internal/domain"

var _ domain.IUsecase = (*Usecase)(nil) // compile-time check

type Usecase struct {
	repo domain.IRepository
}

func New(repo domain.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
