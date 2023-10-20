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
	UserId          int
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

type ApplyRecord struct {
	Id              int
	Detail          string
	CurriculumVitae string
	Works           string
	Skills          string
	Name            string
	Level           int32
	Avatar          string
	CreateAt        time.Time
	UpdateAt        time.Time
	IsPassed        int
	Messages        string
	AuditorId       int
	UserId          int
}

type ReviewApplyRecord struct {
	Id        int
	AuditorId int
	Messages  string
	IsPassed  int
}

type BecomeRecordFilter struct {
	IsPassed *int32
	UserId   *int32
	Id       *int32
	PageNum  *int32
	PageSize *int32
}

type TeacherRepo interface {
	ListAll(ctx context.Context, pageNum, pageSize *int32) ([]*Teacher, error)
	GetById(ctx context.Context, id int) (*Teacher, error)
	Create(ctx context.Context, teacher *Teacher) (*Teacher, error)
	Become(ctx context.Context, applyRecord *ApplyRecord) (*ApplyRecord, error)
	Review(ctx context.Context, reviewApplyRecord *ReviewApplyRecord) (*ApplyRecord, error)
	BecomeRecord(ctx context.Context, becomeRecordFilter *BecomeRecordFilter) ([]*ApplyRecord, error)
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

func (tc *TeacherCase) Become(ctx context.Context, applyRecord *ApplyRecord) (*ApplyRecord, error) {
	return tc.repo.Become(ctx, applyRecord)
}

func (tc *TeacherCase) Review(ctx context.Context, reviewApplyRecord *ReviewApplyRecord) (*ApplyRecord, error) {
	return tc.repo.Review(ctx, reviewApplyRecord)
}

func (tc *TeacherCase) BecomeRecord(ctx context.Context, becomeRecordFilter *BecomeRecordFilter) ([]*ApplyRecord, error) {
	return tc.repo.BecomeRecord(ctx, becomeRecordFilter)
}
