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

type SliderUsecase struct {
	repo SliderRepo
	log  *log.Helper
}

func NewSliderUsecase(repo SliderRepo, logger log.Logger) *SliderUsecase {
	return &SliderUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (su *SliderUsecase) CreateSlider(ctx context.Context, slider *Slider) (*Slider, error) {
	return su.repo.Create(ctx, slider)
}

func (su *SliderUsecase) GetSliderById(ctx context.Context, id int) (*Slider, error) {
	return su.repo.Get(ctx, id)
}

func (su *SliderUsecase) DeleteById(ctx context.Context, id int) (int, error) {
	return su.repo.Delete(ctx, id)
}

func (su *SliderUsecase) UpdateById(ctx context.Context, id int, slider *Slider) (*Slider, error) {
	return su.repo.Update(ctx, id, slider)
}

func (su *SliderUsecase) ListByPriority(ctx context.Context, priority, count int) ([]*Slider, error) {
	return su.repo.ListByPriority(ctx, priority, count)
}
