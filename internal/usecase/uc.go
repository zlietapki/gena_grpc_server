// start name:top
package usecase

import "github.com/zlietapki/gena/internal/domain"

var _ domain.IUsecase = (*Usecase)(nil) // compile-time check

type Usecase struct {
	// start name:uc_struct type:add
	// start name:post_uc_struct
}

type Depends struct {
	// start name:deps type:add
	// start name:post_deps
}

func New(deps Depends) *Usecase {
	return &Usecase{
		// start name:new type:add
		// start name:post_new
	}
}
