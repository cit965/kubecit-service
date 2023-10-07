package service

import (
	"context"
	"errors"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// MostNew 最新好课
func (s *KubecitService) MostNew(ctx context.Context, req *pb.Empty) (*pb.MostNewReply, error) {

	courses, total, err := s.cc.SearchCourse(ctx, &biz.SearchFilterParam{
		PageNum:  GetInt32Ptr(1),
		PageSize: GetInt32Ptr(20),
	})
	if err != nil {
		return nil, err
	}

	result := []*pb.CourseInfo{}
	for _, v := range courses {
		tmp := &pb.CourseInfo{
			Id:         int32(v.Id),
			Level:      pb.CourseLevel(v.Level),
			Name:       v.Name,
			Detail:     v.Detail,
			Cover:      v.Cover,
			Price:      v.Price,
			Tags:       strings.Split(v.Tags, ","),
			CreatedAt:  timestamppb.New(v.CreatedAt),
			UpdatedAt:  timestamppb.New(v.UpdatedAt),
			Status:     pb.CourseStatus(v.Status),
			CategoryId: int32(v.CategoryId),
			People:     102,
			Score:      99,
			Duration:   40,
		}
		result = append(result, tmp)
	}
	return &pb.MostNewReply{List: result, Total: total}, nil
}

// SearchCourse 搜索课程
func (s *KubecitService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.CourseSearchReply, error) {

	courses, total, err := s.cc.SearchCourse(ctx, &biz.SearchFilterParam{
		SecondCategoryId: req.SecondCategory,
		FirstCategoryId:  req.FirstCategory,
		Level:            req.GetLevel(),
		Order:            req.Order,
		PageNum: func() *int32 {
			if req.PageNum == nil {
				return GetInt32Ptr(int32(1))
			} else {
				return req.PageNum
			}
		}(),
		PageSize: func() *int32 {
			if req.PageSize == nil {
				return GetInt32Ptr(int32(20))
			} else {
				return req.PageSize
			}
		}(),
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pb.CourseInfo, 0, 0)
	for _, course := range courses {
		list = append(list, &pb.CourseInfo{
			Id:         int32(course.Id),
			Level:      pb.CourseLevel(course.Level),
			Name:       course.Name,
			Detail:     course.Detail,
			Cover:      course.Cover,
			Price:      course.Price,
			Tags:       strings.Split(course.Tags, ","),
			CreatedAt:  timestamppb.New(course.CreatedAt),
			UpdatedAt:  timestamppb.New(course.UpdatedAt),
			Status:     pb.CourseStatus(course.Status),
			CategoryId: int32(course.CategoryId),
			People:     102,
			Score:      99,
			Duration:   40,
		})
	}
	return &pb.CourseSearchReply{
		List:  list,
		Total: total,
	}, nil
}

// UpdateCourse 更新课程
func (s *KubecitService) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseReply, error) {
	user, err := s.userUseCase.CurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	} else if uint8(user.RoleId) < biz.UserRoleLecturer {
		return nil, errors.New("not enough privileges")
	}
	course := &biz.Course{
		Id:         int(req.GetId()),
		Level:      int32(req.GetLevel().Number()),
		Name:       req.GetName(),
		Detail:     req.GetDetail(),
		Cover:      req.GetCover(),
		Price:      req.GetPrice(),
		Tags:       strings.Join(req.Tags, ","),
		CategoryId: int(req.GetCategoryId()),
	}
	res, err := s.cc.UpdateCourse(ctx, int(req.Id), course)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCourseReply{
		Data: &pb.CourseInfo{
			Id:         int32(res.Id),
			Level:      pb.CourseLevel(res.Level),
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
			People:     102,
			Score:      99,
			Duration:   40,
		}}, nil
}

// ReviewCourse  课程审核
func (s *KubecitService) ReviewCourse(ctx context.Context, req *pb.ReviewCourseRequest) (*pb.ReviewCourseReply, error) {
	user, err := s.userUseCase.CurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	} else if uint8(user.RoleId) < biz.UserRoleSuperAdmin {
		return nil, errors.New("not enough privileges")
	}

	res, err := s.cc.ReviewCourse(ctx, int(req.Id), int32(req.Status))
	if err != nil {
		return nil, err
	}
	return &pb.ReviewCourseReply{
		Data: &pb.CourseInfo{
			Id:         int32(res.Id),
			Level:      pb.CourseLevel(res.Level),
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
			People:     102,
			Score:      99,
			Duration:   40,
		}}, nil
}

func (s *KubecitService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseReply, error) {
	ins := &biz.Course{
		Level:      int32(req.GetLevel().Number()),
		Name:       req.GetName(),
		Detail:     req.GetDetail(),
		Cover:      req.GetCover(),
		Price:      req.GetPrice(),
		Tags:       strings.Join(req.GetTags(), ","),
		CategoryId: int(req.GetCategoryId()),
	}
	res, err := s.cc.CreateCourse(ctx, ins)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCourseReply{Data: &pb.CourseInfo{
		Id:         int32(res.Id),
		Level:      pb.CourseLevel(res.Level),
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       strings.Split(res.Tags, ","),
		Status:     pb.CourseStatus(res.Status),
		CategoryId: int32(res.CategoryId),
		CreatedAt:  timestamppb.New(res.CreatedAt),
		UpdatedAt:  timestamppb.New(res.UpdatedAt),
		People:     102,
		Score:      99,
		Duration:   40,
	}}, nil
}

func (s *KubecitService) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseReply, error) {
	res, err := s.cc.GetCourse(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetCourseReply{
		Data: &pb.CourseInfo{
			Id:         int32(res.Id),
			Level:      pb.CourseLevel(res.Level),
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
			People:     102,
			Score:      99,
			Duration:   40,
		}}, nil
}

// DeleteCourse TODO delete chapter record together
func (s *KubecitService) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseReply, error) {
	count, err := s.cc.DeleteCourse(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCourseReply{
		Count: int32(count)}, nil
}

func (s *KubecitService) CreateChapter(ctx context.Context, req *pb.CreateChapterRequest) (*pb.CreateChapterReply, error) {
	ins := &biz.Chapter{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Sort:        int(req.GetSort()),
		CourseId:    int(req.GetCourseId()),
	}

	res, err := s.cc.CreateChapter(ctx, ins)
	if err != nil {
		return nil, err
	}
	return &pb.CreateChapterReply{Data: &pb.ChapterInfo{
		Id:           int32(res.Id),
		Name:         res.Name,
		ReleasedTime: timestamppb.New(res.ReleasedTime),
		Description:  res.Description,
		Sort:         int32(res.Sort),
		CourseId:     int32(res.CourseId),
	}}, nil
}

func (s *KubecitService) DeleteChapter(ctx context.Context, req *pb.DeleteChapterRequest) (*pb.DeleteChapterReply, error) {
	res, err := s.cc.DeleteChapter(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteChapterReply{
		Count: int32(res),
	}, nil
}

func (s *KubecitService) UpdateChapter(ctx context.Context, req *pb.UpdateChapterRequest) (*pb.UpdateChapterReply, error) {
	ins := &biz.Chapter{
		Id:          int(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Sort:        int(req.GetSort()),
		CourseId:    int(req.GetCourseId()),
	}
	res, err := s.cc.UpdateChapter(ctx, int(req.Id), ins)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateChapterReply{Data: &pb.ChapterInfo{
		Id:           int32(res.Id),
		Name:         res.Name,
		ReleasedTime: timestamppb.New(res.ReleasedTime),
		Description:  res.Description,
		Sort:         int32(res.Sort),
		CourseId:     int32(res.CourseId),
	}}, nil
}

func (s *KubecitService) CreateLesson(ctx context.Context, req *pb.CreateLessonRequest) (*pb.CreateLessonReply, error) {
	ins := &biz.Lesson{
		Name:          req.GetName(),
		Sort:          int(req.GetSort()),
		Type:          int(req.GetType()),
		StoragePath:   req.GetStoragePath(),
		Source:        req.GetSource(),
		Courseware:    req.GetCourseware(),
		IsFreePreview: int(req.GetIsFreePreview()),
		ChapterId:     int(req.GetChapterId()),
	}
	res, err := s.cc.CreateLesson(ctx, ins)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLessonReply{Data: &pb.LessonInfo{
		Name:          res.Name,
		Sort:          int32(res.Sort),
		Type:          pb.LessonType(res.Type),
		StoragePath:   res.StoragePath,
		Source:        res.Source,
		Courseware:    res.Courseware,
		IsFreePreview: pb.FreePreview(res.IsFreePreview),
		ChapterId:     int32(res.ChapterId),
		Id:            int32(res.Id),
		ReleasedTime:  timestamppb.New(res.ReleasedTime),
	}}, nil
}
func (s *KubecitService) ListLessons(ctx context.Context, req *pb.ListChapterLessonsWithCourIdRequest) (*pb.ListChapterLessonsWithCourIdReply, error) {
	chapterLessons, err := s.cc.ListChapterLessonsWithCourseId(ctx, int(req.GetCourseId()))
	if err != nil {
		return nil, err
	}
	data := make([]*pb.ChapterLessonsInfo, 0, len(chapterLessons))
	for _, chapterLesson := range chapterLessons {
		lessonData := make([]*pb.LessonInfo, 0, len(chapterLesson.Lessons))
		for _, lesson := range chapterLesson.Lessons {
			lessonData = append(lessonData, &pb.LessonInfo{
				Name:          lesson.Name,
				Sort:          int32(lesson.Sort),
				Type:          pb.LessonType(lesson.Type),
				StoragePath:   lesson.StoragePath,
				Source:        lesson.Source,
				Courseware:    lesson.Courseware,
				IsFreePreview: pb.FreePreview(lesson.IsFreePreview),
				ChapterId:     int32(lesson.ChapterId),
				Id:            int32(lesson.Id),
				ReleasedTime:  timestamppb.New(lesson.ReleasedTime),
			})
		}
		data = append(data, &pb.ChapterLessonsInfo{
			Chapter: &pb.ChapterInfo{
				Id:           int32(chapterLesson.Id),
				Name:         chapterLesson.Name,
				ReleasedTime: timestamppb.New(chapterLesson.ReleasedTime),
				Description:  chapterLesson.Description,
				Sort:         int32(chapterLesson.Sort),
				CourseId:     int32(chapterLesson.CourseId),
			},
			Lessons: lessonData,
		})
	}

	return &pb.ListChapterLessonsWithCourIdReply{Data: data}, nil
}
func (s *KubecitService) UpdateLesson(ctx context.Context, req *pb.UpdateLessonRequest) (*pb.UpdateLessonReply, error) {
	ins := &biz.Lesson{
		Name:          req.GetName(),
		Sort:          int(req.GetSort()),
		Type:          int(req.GetType()),
		StoragePath:   req.GetStoragePath(),
		Source:        req.GetSource(),
		Courseware:    req.GetCourseware(),
		IsFreePreview: int(req.GetIsFreePreview()),
		ChapterId:     int(req.GetChapterId()),
	}
	res, err := s.cc.UpdateLesson(ctx, int(req.GetLessonId()), ins)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateLessonReply{
		Data: &pb.LessonInfo{
			Name:          res.Name,
			Sort:          int32(res.Sort),
			Type:          pb.LessonType(res.Type),
			StoragePath:   res.StoragePath,
			Source:        res.Source,
			Courseware:    res.Courseware,
			IsFreePreview: pb.FreePreview(res.IsFreePreview),
			ChapterId:     int32(res.ChapterId),
			Id:            int32(res.Id),
			ReleasedTime:  timestamppb.New(res.ReleasedTime),
		},
	}, nil
}
func (s *KubecitService) DeleteLesson(ctx context.Context, req *pb.DeleteLessonRequest) (*pb.DeleteLessonReply, error) {
	count, err := s.cc.DeleteLesson(ctx, int(req.GetLessonId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteLessonReply{Count: int32(count)}, nil
}

func GetInt32Ptr(src int32) *int32 {
	return &src
}
