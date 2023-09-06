package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Slider struct {
	Id        int
	Title     string
	Content   string
	ImageLink string
	CreateAt  time.Time
	UpdateAt  time.Time
	IsValid   bool
	Priority  int
}

type SliderRepo interface {
	Create(context.Context, *Slider) (*Slider, error)
	Get(ctx context.Context, id int) (*Slider, error)
	List(ctx context.Context) ([]*Slider, error)
	Delete(ctx context.Context, id int) (int, error)
	Update(ctx context.Context, id int, slider *Slider) (*Slider, error)
	ListByPriority(ctx context.Context, priority int, count int) ([]*Slider, error)
}

type SystemUsecase struct {
	repo SliderRepo
	log  *log.Helper
}

func NewSystemUsecase(repo SliderRepo, logger log.Logger) *SystemUsecase {
	return &SystemUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (su *SystemUsecase) CreateSlider(ctx context.Context, slider *Slider) (*Slider, error) {
	return su.repo.Create(ctx, slider)
}

func (su *SystemUsecase) GetSliderById(ctx context.Context, id int) (*Slider, error) {
	return su.repo.Get(ctx, id)
}

func (su *SystemUsecase) DeleteById(ctx context.Context, id int) (int, error) {
	return su.repo.Delete(ctx, id)
}

func (su *SystemUsecase) UpdateById(ctx context.Context, id int, slider *Slider) (*Slider, error) {
	return su.repo.Update(ctx, id, slider)
}

func (su *SystemUsecase) ListByPriority(ctx context.Context, priority, count int) ([]*Slider, error) {
	return su.repo.ListByPriority(ctx, priority, count)
}
