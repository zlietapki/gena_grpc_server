package usecase

import (
	"context"
	"fmt"
)

func (u Usecase) Hello(ctx context.Context, name string) (string, error) {
	res := fmt.Sprintf("Hello: %s", name)

	return res, nil
}
