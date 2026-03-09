// start name:top
package domain

//start name:imports type:merge
import "context"

// start name:usecase
type IUsecase interface {
	//start name:methods type:merge
	Hello(ctx context.Context, name string) (string, error)
	//start name:post_methods
}
