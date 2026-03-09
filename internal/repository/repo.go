// start name:top
package repository

// start name:imports type:merge
import "github.com/zlietapki/boilerplate/internal/domain"

// start name:post_imports

var _ domain.IRepository = (*Repo)(nil) // compile-time check

type Repo struct {
	//start name:repo_struct type:merge
	// start name:post_repo_struct
}

func New() *Repo {
	return &Repo{
		// start name:repo_init type:merge
		// start name:post_repo_init
	}
}
