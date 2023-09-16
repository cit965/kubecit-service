package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Teacher struct {
	Id              int
	Detail          string
	CurriculumVitae string
	Works           string
	Skills          string
	Avator          string
	CreateAt        time.Time
	UpdateAt        time.Time
}

type TeacherRepo interface {
	ListAll(ctx context.Context) ([]*Teacher, error)
	GetById(ctx context.Context, id int) (*Teacher, error)
}

type TeacherCase struct {
	repo TeacherRepo
	log  *log.Helper
}

func NewTeacherCase(repo TeacherRepo, logger log.Logger) *TeacherCase {
	return &TeacherCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TeacherCase) ListTeachers(ctx context.Context) ([]*Teacher, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *TeacherCase) GetTeacher(ctx context.Context, id int) (*Teacher, error) {
	return uc.repo.GetById(ctx, id)
}
