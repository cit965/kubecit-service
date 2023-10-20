package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kubecit-service/ent"
	"kubecit-service/ent/applyrecord"
	"kubecit-service/ent/teacher"
	"kubecit-service/internal/biz"
	"kubecit-service/internal/pkg/common"
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
	limit, offset := common.ConvertPageSize(pageNum, pageSize)
	teachers, err := cq.Limit(limit).Offset(offset).All(ctx)

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
	teacherObj, err := t.data.db.Teacher.Query().Where(teacher.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, errors.NotFound(err.Error(), "未找到讲师")
	}
	return &biz.Teacher{
		Id:              teacherObj.ID,
		Detail:          teacherObj.Detail,
		CurriculumVitae: teacherObj.CurriculumVitae,
		Works:           teacherObj.Works,
		Skills:          teacherObj.Skills,
		Avator:          teacherObj.Avator,
		Name:            teacherObj.Name,
		Level:           int32(teacherObj.Level),
		CreateAt:        teacherObj.CreateAt,
		UpdateAt:        teacherObj.UpdateAt,
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

func (t *teacherRepo) Become(ctx context.Context, applyRecord *biz.ApplyRecord) (*biz.ApplyRecord, error) {
	tc, err := t.data.db.ApplyRecord.Create().SetDetail(applyRecord.Detail).SetCurriculumVitae(applyRecord.CurriculumVitae).
		SetWorks(applyRecord.Works).SetSkills(applyRecord.Skills).SetName(applyRecord.Name).
		SetLevel(int(applyRecord.Level)).SetAvatar(applyRecord.Avatar).SetIsPassed(2).SetUserID(applyRecord.UserId).Save(ctx)
	if err != nil {
		t.log.Errorf("course repo update error: %v\n", err)
		return nil, err
	}

	return &biz.ApplyRecord{
		Id:              tc.ID,
		Detail:          tc.Detail,
		CurriculumVitae: tc.CurriculumVitae,
		Works:           tc.Works,
		Skills:          tc.Skills,
		Name:            tc.Name,
		Level:           int32(tc.Level),
		Avatar:          tc.Avatar,
		CreateAt:        tc.CreateAt,
		UpdateAt:        tc.UpdateAt,
		IsPassed:        int(tc.IsPassed),
		//Messages:        tc.Messages,
		//AuditorId:       tc.AuditorID,
		UserId: tc.UserID,
	}, nil
}

func (t *teacherRepo) Review(ctx context.Context, reviewApplyRecord *biz.ReviewApplyRecord) (*biz.ApplyRecord, error) {
	result, err := t.data.WithResultTx(ctx, func(tx *ent.Tx) (interface{}, error) {
		ar, err := t.review(ctx, reviewApplyRecord)
		if err != nil {
			return nil, err
		}
		return ar, nil
	})
	if err != nil {
		return nil, err
	}
	if ar, ok := result.(*biz.ApplyRecord); ok {
		return ar, nil
	}
	return nil, err
}

func (t *teacherRepo) review(ctx context.Context, reviewApplyRecord *biz.ReviewApplyRecord) (*biz.ApplyRecord, error) {
	ar, err := t.data.db.ApplyRecord.UpdateOneID(reviewApplyRecord.Id).SetIsPassed(int8(reviewApplyRecord.IsPassed)).SetMessages(reviewApplyRecord.Messages).
		SetAuditorID(reviewApplyRecord.AuditorId).Save(ctx)
	if err != nil {
		t.log.Errorf("teacher repo update error: %v\n", err)
		return nil, err
	}
	if reviewApplyRecord.IsPassed == 1 {
		tc, terr := t.data.db.Teacher.Create().SetDetail(ar.Detail).SetCurriculumVitae(ar.CurriculumVitae).SetWorks(ar.Works).
			SetSkills(ar.Skills).SetAvator(ar.Avatar).SetName(ar.Name).SetLevel(ar.Level).SetUserID(ar.UserID).Save(ctx)
		if terr != nil {
			t.log.Errorf("teacher repo update error: %v\n", terr)
			return nil, terr
		}
		_, err = t.data.db.User.UpdateOneID(tc.UserID).SetRoleID(3).Save(ctx)
		if err != nil {
			t.log.Errorf("teacher repo update user role error: %v\n", terr)
			return nil, terr
		}
	}
	return &biz.ApplyRecord{
		Id:              ar.ID,
		Detail:          ar.Detail,
		CurriculumVitae: ar.CurriculumVitae,
		Works:           ar.Works,
		Skills:          ar.Skills,
		Name:            ar.Name,
		Level:           int32(ar.Level),
		Avatar:          ar.Avatar,
		CreateAt:        ar.CreateAt,
		UpdateAt:        ar.UpdateAt,
		IsPassed:        int(ar.IsPassed),
		Messages:        ar.Messages,
		AuditorId:       ar.AuditorID,
		UserId:          ar.UserID,
	}, nil
}

func (t *teacherRepo) BecomeRecord(ctx context.Context, becomeRecordFilter *biz.BecomeRecordFilter) ([]*biz.ApplyRecord, error) {
	q := t.data.db.ApplyRecord.Query()
	if becomeRecordFilter.Id != nil {
		q.Where(applyrecord.IDEQ(int(*becomeRecordFilter.Id)))
	}
	if becomeRecordFilter.IsPassed != nil {
		q.Where(applyrecord.IsPassedEQ(int8(*becomeRecordFilter.IsPassed)))
	}
	if becomeRecordFilter.UserId != nil {
		q.Where(applyrecord.UserIDEQ(int(*becomeRecordFilter.UserId)))
	}

	*becomeRecordFilter.PageNum--
	q.Offset(int(*becomeRecordFilter.PageNum) * int(*becomeRecordFilter.PageSize))
	q.Limit(int(*becomeRecordFilter.PageSize))

	ars, err := q.All(ctx)
	if err != nil {
		t.log.Errorf("BecomeRecord repo get error: %v\n", err)
		return nil, err
	}
	var result []*biz.ApplyRecord
	for _, ar := range ars {
		result = append(result, &biz.ApplyRecord{
			Id:              ar.ID,
			Detail:          ar.Detail,
			CurriculumVitae: ar.CurriculumVitae,
			Works:           ar.Works,
			Skills:          ar.Skills,
			Name:            ar.Name,
			Level:           int32(ar.Level),
			Avatar:          ar.Avatar,
			CreateAt:        ar.CreateAt,
			UpdateAt:        ar.UpdateAt,
			IsPassed:        int(ar.IsPassed),
			Messages:        ar.Messages,
			AuditorId:       ar.AuditorID,
			UserId:          ar.UserID,
		})
	}
	return result, nil
}
