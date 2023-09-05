package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Category is a Category model.
type Category struct {
	CategoryName string
	Id           int32
	ParentId     int32
	Level        string
	Status       string
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	ListAll(ctx context.Context) ([]*Category, error)

	ListFirstCategories(ctx context.Context) ([]*Category, error)
}

// CourseUsecase is a Category usecase.
type CourseUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCourseUsecase new a Category usecase.
func NewCourseUsecase(repo CategoryRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CourseUsecase) ListCategory(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *CourseUsecase) ListFirstCategory(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListFirstCategories(ctx)
}
