package biz

import (
	"context"

	v1 "kubecit-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

type Category struct {
	CategoryName string
	Id           string
}

type CategoryRepo interface {
	ListAll(ctx context.Context) ([]*Category, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo         GreeterRepo
	categoryRepo CategoryRepo
	log          *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, categoryRepo CategoryRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, categoryRepo: categoryRepo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) ListCategory(ctx context.Context) ([]*Category, error) {
	return uc.categoryRepo.ListAll(ctx)
}
