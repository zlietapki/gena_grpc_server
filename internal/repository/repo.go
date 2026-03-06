package repository

import "github.com/zlietapki/microboiler_grpc_server/internal/domain"

var _ domain.IRepository = (*Repo)(nil) // compile-time check

type Repo struct{}

func New() *Repo {
	return &Repo{}
}
