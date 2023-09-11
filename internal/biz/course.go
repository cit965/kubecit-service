package biz

import (
	"context"
	"time"

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

type Course struct {
	Id         int
	Level      int32
	Name       string
	Detail     string
	Cover      string
	Price      float32
	Tags       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     int32
	CategoryId int
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	ListAll(ctx context.Context) ([]*Category, error)

	ListFirstCategories(ctx context.Context) ([]*Category, error)
}

// CourseRepo is a Course repo.
type CourseRepo interface {
	SearchCourse(ctx context.Context, pageNum, pageSize int, categoryId *int32, level *int32, reverse *bool) ([]*Course, error)
	UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error)
	ReviewCourse(ctx context.Context, id int, status int32) (*Course, error)
}

// CourseUsecase is a Category usecase.
type CourseUsecase struct {
	repo       CategoryRepo
	courseRepo CourseRepo
	log        *log.Helper
}

// NewCourseUsecase new a Category usecase.
func NewCourseUsecase(repo CategoryRepo, courseRepo CourseRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{
		repo:       repo,
		courseRepo: courseRepo,
		log:        log.NewHelper(logger),
	}
}

func (uc *CourseUsecase) ListCategory(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *CourseUsecase) ListFirstCategory(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListFirstCategories(ctx)
}

func (uc *CourseUsecase) SearchCourse(ctx context.Context, pageNum, pageSize int, categoryId *int32, level *int32, reverse *bool) ([]*Course, error) {
	return uc.courseRepo.SearchCourse(ctx, pageNum, pageSize, categoryId, level, reverse)
}

func (uc *CourseUsecase) UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error) {
	return uc.courseRepo.UpdateCourse(ctx, id, course)
}

func (uc *CourseUsecase) ReviewCourse(ctx context.Context, id int, status int32) (*Course, error) {
	return uc.courseRepo.ReviewCourse(ctx, id, status)
}
