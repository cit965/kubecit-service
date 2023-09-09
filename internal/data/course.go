package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent"
	"kubecit-service/ent/course"
	"kubecit-service/internal/biz"
)

type courseRepo struct {
	data *Data
	log  *log.Helper
}

func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &courseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c courseRepo) SearchCourse(ctx context.Context, pageNum, pageSize int, categoryId *int32, level *int32, reverse *bool) ([]*biz.Course, error) {
	cq := c.data.db.Course.Query()
	if categoryId != nil {
		cq.Where(course.CategoryIDEQ(int(*categoryId)))
	}
	if level != nil {
		cq.Where(course.LevelEQ(*level))
	}
	if reverse != nil {
		if !*reverse {
			cq.Order(ent.Asc(course.FieldCreatedAt))
		} else {
			cq.Order(ent.Desc(course.FieldCreatedAt))
		}
	}

	result, err := cq.Offset(0).Limit(pageSize).All(ctx)
	if err != nil {
		c.log.Errorf("search course errorf: %v\n", err)
		return nil, err
	}
	courses := make([]*biz.Course, 0, len(result))
	for _, v := range result {
		courses = append(courses, &biz.Course{
			Id:         v.ID,
			Level:      v.Level,
			Name:       v.Name,
			Detail:     v.Detail,
			Cover:      v.Cover,
			Price:      v.Price,
			Tags:       v.Tags,
			CreatedAt:  v.CreatedAt,
			Status:     v.Status,
			CategoryId: v.CategoryID,
		})
	}
	return courses, nil
}
