package usecase

import (
	"context"
	"fmt"
)

func (u UC) Hello(ctx context.Context, name string) (string, error) {
	res := fmt.Sprintf("Hello: %s\n", name)

	//u.repo.

	return res, nil
}
