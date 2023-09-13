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
	Level        int
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
	ListByLevel(ctx context.Context, level int32) ([]*Category, error)
	Create(context.Context, *Category) error
	ListSubCategories(ctx context.Context, parentID int32) ([]*Category, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, id int, name string) error
}

// CourseRepo is a Course repo.
type CourseRepo interface {
	SearchCourse(ctx context.Context, pageNum, pageSize int, categoryIds []int, level int32, reverse *bool) ([]*Course, error)
	UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error)
	ReviewCourse(ctx context.Context, id int, status int32) (*Course, error)
	CreateCourse(ctx context.Context, course *Course) (*Course, error)
	GetCourse(ctx context.Context, id int) (*Course, error)
	DeleteCourse(ctx context.Context, id int) (int, error)
}

// CourseUsecase is a Category usecase.
type CourseUsecase struct {
	repo       CategoryRepo
	courseRepo CourseRepo

	log *log.Helper
}

// NewCourseUsecase new a Category usecase.
func NewCourseUsecase(repo CategoryRepo, courseRepo CourseRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{
		repo:       repo,
		courseRepo: courseRepo,
		log:        log.NewHelper(logger),
	}
}

func (uc *CourseUsecase) ListCategory(ctx context.Context, level int32) ([]*Category, error) {
	return uc.repo.ListByLevel(ctx, level)
}

func (uc *CourseUsecase) CreateCategory(ctx context.Context, category *Category) error {
	return uc.repo.Create(ctx, category)
}

func (uc *CourseUsecase) DeleteCategory(ctx context.Context, id int32) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *CourseUsecase) UpdateCategory(ctx context.Context, id int, name string) error {
	return uc.repo.Update(ctx, id, name)
}

type SearchFilterParam struct {
	SecondCategoryId int32
	FirstCategoryId  int32
	Level            int32
	Reverse          bool
}

func (uc *CourseUsecase) SearchCourse(ctx context.Context, pageNum, pageSize int, filter *SearchFilterParam) ([]*Course, error) {

	var categoryIds []int
	if filter.SecondCategoryId == 0 {
		if filter.FirstCategoryId != 0 {
			subCategories, err := uc.repo.ListSubCategories(ctx, filter.FirstCategoryId)
			if err != nil {
				return nil, err
			}
			for _, v := range subCategories {
				categoryIds = append(categoryIds, int(v.Id))
			}
		}

	} else {
		categoryIds = append(categoryIds, int(filter.SecondCategoryId))
	}

	return uc.courseRepo.SearchCourse(ctx, pageNum, pageSize, categoryIds, filter.Level, &filter.Reverse)
}

func (uc *CourseUsecase) UpdateCourse(ctx context.Context, id int, course *Course) (*Course, error) {
	return uc.courseRepo.UpdateCourse(ctx, id, course)
}

func (uc *CourseUsecase) ReviewCourse(ctx context.Context, id int, status int32) (*Course, error) {
	return uc.courseRepo.ReviewCourse(ctx, id, status)
}

func (uc *CourseUsecase) CreateCourse(ctx context.Context, course *Course) (*Course, error) {
	return uc.courseRepo.CreateCourse(ctx, course)
}

func (uc *CourseUsecase) GetCourse(ctx context.Context, id int) (*Course, error) {
	return uc.courseRepo.GetCourse(ctx, id)
}

func (uc *CourseUsecase) DeleteCourse(ctx context.Context, id int) (int, error) {
	return uc.courseRepo.DeleteCourse(ctx, id)
}
