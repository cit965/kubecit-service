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
	Name            string
	Level           int32
	CreateAt        time.Time
	UpdateAt        time.Time
}

type RecommendedLecturer struct {
	Id          int
	Name        string
	Avatar      string
	Level       string
	Title       string
	Brief       string
	Achievement *Achievement
}

type Achievement struct {
	Students   int32
	Fans       int32
	ClassHours int32
	Praises    int32
}

type TeacherRepo interface {
	ListAll(ctx context.Context, pageNum, pageSize *int32) ([]*Teacher, error)
	GetById(ctx context.Context, id int) (*Teacher, error)
	Create(ctx context.Context, teacher *Teacher) (*Teacher, error)
}

type TeacherCase struct {
	repo TeacherRepo
	log  *log.Helper
}

func NewTeacherCase(repo TeacherRepo, logger log.Logger) *TeacherCase {
	return &TeacherCase{repo: repo, log: log.NewHelper(logger)}
}

func (tc *TeacherCase) ListTeachers(ctx context.Context, pageNum, pageSize *int32) ([]*Teacher, error) {
	return tc.repo.ListAll(ctx, pageNum, pageSize)
}

func (tc *TeacherCase) GetTeacher(ctx context.Context, id int) (*Teacher, error) {
	return tc.repo.GetById(ctx, id)
}

func (tc *TeacherCase) CreateTeacher(ctx context.Context, teacher *Teacher) (*Teacher, error) {
	return tc.repo.Create(ctx, teacher)
}
