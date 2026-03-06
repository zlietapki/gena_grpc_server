package domain

import "context"

type IUsecase interface {
	Hello(ctx context.Context, name string) (string, error)
}
