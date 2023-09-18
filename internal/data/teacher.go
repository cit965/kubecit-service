package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kubecit-service/ent/teacher"
	"kubecit-service/internal/biz"
)

type teacherRepo struct {
	data *Data
	log  *log.Helper
}

func NewTeacherRepo(data *Data, logger log.Logger) biz.TeacherRepo {
	return &teacherRepo{data: data, log: log.NewHelper(logger)}
}

func (t *teacherRepo) ListAll(ctx context.Context, pageNum, pageSize *int32) ([]*biz.Teacher, error) {
	cq := t.data.db.Teacher.Query()
	if pageNum != nil {
		*pageNum--
		cq.Offset(int(*pageNum))
	} else {
		cq.Offset(0)
	}
	if pageSize != nil {
		cq.Limit(int(*pageSize))
	} else {
		cq.Limit(20)
	}

	teachers, err := cq.All(ctx)

	if err != nil {
		return nil, errors.BadRequest(err.Error(), "获取讲师列表失败！")
	}
	teacherResult := make([]*biz.Teacher, 0)
	err = copier.Copy(&teacherResult, &teachers)
	if err != nil {
		return nil, errors.BadRequest(err.Error(), "copier 操作失败")
	}

	return teacherResult, nil
}

func (t *teacherRepo) GetById(ctx context.Context, id int) (*biz.Teacher, error) {
	teacher, err := t.data.db.Teacher.Query().Where(teacher.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, errors.NotFound(err.Error(), "未找到讲师")
	}
	return &biz.Teacher{
		Id:              teacher.ID,
		Detail:          teacher.Detail,
		CurriculumVitae: teacher.CurriculumVitae,
		Works:           teacher.Works,
		Skills:          teacher.Skills,
		Avator:          teacher.Avator,
		Name:            teacher.Name,
		Level:           int32(teacher.Level),
		CreateAt:        teacher.CreateAt,
		UpdateAt:        teacher.UpdateAt,
	}, nil
}

func (t *teacherRepo) Create(ctx context.Context, teacher *biz.Teacher) (*biz.Teacher, error) {
	tc, err := t.data.db.Teacher.Create().SetDetail(teacher.Detail).SetCurriculumVitae(teacher.CurriculumVitae).SetWorks(teacher.Works).
		SetSkills(teacher.Skills).SetAvator(teacher.Avator).SetName(teacher.Name).SetLevel(int(teacher.Level)).Save(ctx)
	if err != nil {
		return nil, errors.BadRequest(err.Error(), "讲师创建失败")
	}
	teacherModel := &biz.Teacher{
		Id:              tc.ID,
		Detail:          tc.Detail,
		CurriculumVitae: tc.CurriculumVitae,
		Works:           tc.Works,
		Skills:          tc.Skills,
		Avator:          tc.Avator,
		Name:            tc.Name,
		Level:           int32(tc.Level),
		CreateAt:        tc.CreateAt,
		UpdateAt:        tc.UpdateAt,
	}
	return teacherModel, nil
}
