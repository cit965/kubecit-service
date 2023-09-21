package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/ent"
	"kubecit-service/ent/chapter"
	"kubecit-service/ent/course"
	"kubecit-service/ent/lesson"
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

func (c *courseRepo) SearchCourse(ctx context.Context, pageNum, pageSize *int32, categories []int, level pb.CourseLevel, order *int32) ([]*biz.Course, int32, error) {
	cq := c.data.db.Course.Query()
	if len(categories) != 0 {
		cq.Where(course.CategoryIDIn(categories...))
	}

	if level != pb.CourseLevel_LEVEL_UNKNOWN {
		cq.Where(course.LevelEQ(int32(level)))
	}
	if order != nil {
		if *order == 1 {
			// 最新排序
			cq.Order(ent.Asc(course.FieldCreatedAt))
		} else if *order == 2 {
			// 综合排序
			cq.Order(ent.Desc(course.FieldCreatedAt))
		}
	}
	total, err := cq.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
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

	result, err := cq.All(ctx)
	if err != nil {
		c.log.Errorf("search course errorf: %v\n", err)
		return nil, 0, err
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
			UpdatedAt:  v.UpdatedAt,
			Status:     v.Status,
			CategoryId: v.CategoryID,
		})
	}
	return courses, int32(total), nil
}

func (c *courseRepo) GetCourse(ctx context.Context, id int) (*biz.Course, error) {
	//res, err := c.data.db.Course.Query().Where(course.IDEQ(id)).Only(ctx)
	res, err := c.data.db.Course.Query().Where(course.IDEQ(id)).Only(ctx)
	if err != nil {
		c.log.Errorf("course repo get error: %v\n", err)
		return nil, err
	}

	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) UpdateCourse(ctx context.Context, id int, ins *biz.Course) (*biz.Course, error) {
	res, err := c.data.db.Course.UpdateOneID(id).SetLevel(ins.Level).SetName(ins.Name).SetDetail(ins.Detail).SetCover(ins.Cover).
		SetPrice(ins.Price).SetTags(ins.Tags).SetCategoryID(ins.CategoryId).Save(ctx)
	if err != nil {
		c.log.Errorf("course repo update error: %v\n", err)
		return nil, err
	}
	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) ReviewCourse(ctx context.Context, id int, status int32) (*biz.Course, error) {
	res, err := c.data.db.Course.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		c.log.Errorf("course repo review error: %v\n", err)
		return nil, err
	}
	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) CreateCourse(ctx context.Context, course *biz.Course) (*biz.Course, error) {
	res, err := c.data.db.Course.Create().SetLevel(course.Level).SetName(course.Name).SetDetail(course.Detail).SetCover(course.Cover).
		SetPrice(course.Price).SetTags(course.Tags).SetStatus(course.Status).SetCategoryID(course.CategoryId).Save(ctx)
	if err != nil {
		c.log.Errorf("course repo create error: %v\n", err)
		return nil, err
	}
	return &biz.Course{
		Id:         res.ID,
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       res.Tags,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Status:     res.Status,
		CategoryId: res.CategoryID,
	}, nil
}

func (c *courseRepo) DeleteCourse(ctx context.Context, id int) (int, error) {
	res, err := c.data.db.Course.Delete().Where(course.IDEQ(id)).Exec(ctx)
	if err != nil {
		c.log.Errorf("course repo delete error: %v\n", err)
		return 0, err
	}
	return res, nil
}

func (c *courseRepo) CreateChapter(ctx context.Context, chapter *biz.Chapter) (*biz.Chapter, error) {
	res, err := c.data.db.Chapter.Create().SetName(chapter.Name).SetDescription(chapter.Description).SetSort(chapter.Sort).SetCourseID(chapter.CourseId).Save(ctx)
	if err != nil {
		c.log.Errorf("chapter repo create error: %v\n", err)
		return nil, err
	}
	return &biz.Chapter{
		Id:           res.ID,
		Name:         res.Name,
		ReleasedTime: res.ReleasedTime,
		Description:  res.Description,
		Sort:         res.Sort,
		CourseId:     res.CourseID,
	}, nil
}

func (c *courseRepo) DeleteChapter(ctx context.Context, id int) (int, error) {
	count, err := c.data.db.Chapter.Delete().Where(chapter.IDEQ(id)).Exec(ctx)
	if err != nil {
		c.log.Errorf("chapter repo delete error: %v\n", err)
		return 0, err
	}
	return count, nil
}

func (c *courseRepo) ListChapters(ctx context.Context, courseId int) ([]*biz.Chapter, error) {
	chapters, err := c.data.db.Debug().Chapter.Query().Where(chapter.CourseIDEQ(courseId)).All(ctx)
	fmt.Printf("%+v\n", chapters)
	if err != nil {
		c.log.Errorf("chapter repo get error: %v\n", err)
		return nil, err
	}
	res := make([]*biz.Chapter, 0, len(chapters))
	for _, ins := range chapters {
		res = append(res, &biz.Chapter{
			Id:           ins.ID,
			Name:         ins.Name,
			ReleasedTime: ins.ReleasedTime,
			Description:  ins.Description,
			Sort:         ins.Sort,
			CourseId:     ins.CourseID,
		})
	}
	return res, nil
}

func (c *courseRepo) UpdateChapter(ctx context.Context, id int, ins *biz.Chapter) (*biz.Chapter, error) {
	res, err := c.data.db.Chapter.UpdateOneID(id).SetName(ins.Name).SetDescription(ins.Description).SetSort(ins.Sort).SetCourseID(ins.CourseId).Save(ctx)
	if err != nil {
		c.log.Errorf("chapter repo get error: %v\n", err)
		return nil, err
	}
	return &biz.Chapter{
		Id:           res.ID,
		Name:         res.Name,
		ReleasedTime: res.ReleasedTime,
		Description:  res.Description,
		Sort:         res.Sort,
		CourseId:     res.CourseID,
	}, nil
}

func (c *courseRepo) CreateLesson(ctx context.Context, lesson *biz.Lesson) (*biz.Lesson, error) {
	res, err := c.data.db.Lesson.Create().SetName(lesson.Name).SetSort(lesson.Sort).
		SetType(lesson.Type).SetStoragePath(lesson.StoragePath).SetSource(lesson.Source).SetCourseware(lesson.Courseware).
		SetIsFreePreview(lesson.IsFreePreview).SetChapterID(lesson.ChapterId).Save(ctx)
	if err != nil {
		c.log.Errorf("lesson repo create error: %v\n", err)
		return nil, err
	}
	return &biz.Lesson{
		Id:            res.ID,
		Name:          res.Name,
		ReleasedTime:  res.ReleasedTime,
		Sort:          res.Sort,
		Type:          res.Type,
		StoragePath:   res.StoragePath,
		Source:        res.Source,
		Courseware:    res.Courseware,
		IsFreePreview: res.IsFreePreview,
		ChapterId:     res.ChapterID,
	}, nil
}

func (c *courseRepo) DeleteLesson(ctx context.Context, id int) (int, error) {
	count, err := c.data.db.Lesson.Delete().Where(lesson.IDEQ(id)).Exec(ctx)
	if err != nil {
		c.log.Errorf("lesson repo delete error: %v\n", err)
		return 0, err
	}
	return count, nil
}

func (c *courseRepo) ListLessons(ctx context.Context, chapterId int) ([]*biz.Lesson, error) {
	lessons, err := c.data.db.Lesson.Query().Where(lesson.ChapterIDEQ(chapterId)).All(ctx)
	if err != nil {
		c.log.Errorf("lesson repo list error: %v\n", err)
		return nil, err
	}
	res := make([]*biz.Lesson, 0, len(lessons))
	for _, ins := range lessons {
		res = append(res, &biz.Lesson{
			Id:            ins.ID,
			Name:          ins.Name,
			ReleasedTime:  ins.ReleasedTime,
			Sort:          ins.Sort,
			Type:          ins.Type,
			StoragePath:   ins.StoragePath,
			Source:        ins.Source,
			Courseware:    ins.Courseware,
			IsFreePreview: ins.IsFreePreview,
			ChapterId:     ins.ChapterID,
		})
	}
	return res, nil
}
func (c *courseRepo) UpdateLesson(ctx context.Context, id int, lesson *biz.Lesson) (*biz.Lesson, error) {
	res, err := c.data.db.Lesson.UpdateOneID(id).SetName(lesson.Name).SetSort(lesson.Sort).
		SetType(lesson.Type).SetStoragePath(lesson.StoragePath).SetSource(lesson.Source).SetCourseware(lesson.Courseware).
		SetIsFreePreview(lesson.IsFreePreview).SetChapterID(lesson.ChapterId).Save(ctx)
	if err != nil {
		c.log.Errorf("lesson repo update error: %v\n", err)
		return nil, err
	}
	return &biz.Lesson{
		Id:            res.ID,
		Name:          res.Name,
		ReleasedTime:  res.ReleasedTime,
		Sort:          res.Sort,
		Type:          res.Type,
		StoragePath:   res.StoragePath,
		Source:        res.Source,
		Courseware:    res.Courseware,
		IsFreePreview: res.IsFreePreview,
		ChapterId:     res.ChapterID,
	}, nil
}
