package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Category is a Category model.
type Category struct {
	CategoryName string
	Id           string
	ParentId     string
	Level        string
	Status       string
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	ListAll(ctx context.Context) ([]*Category, error)
}

// CategoryUsecase is a Category usecase.
type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCategoryUsecase new a Category usecase.
func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CategoryUsecase) ListCategory(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListAll(ctx)
}
