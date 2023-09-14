package service

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
	"strings"
)

// MostNew 最新好课
func (s *KubecitService) MostNew(ctx context.Context, req *pb.Empty) (*pb.MostNewReply, error) {

	courses, total, err := s.cc.SearchCourse(ctx, &biz.SearchFilterParam{})
	if err != nil {
		return nil, err
	}

	result := []*pb.CourseInfo{}
	for _, v := range courses {
		tmp := &pb.CourseInfo{
			Id:         int32(v.Id),
			Level:      v.Level,
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
		}
		result = append(result, tmp)
	}
	return &pb.MostNewReply{List: result, Total: total}, nil
}

// TagsList 课程标签
func (s *KubecitService) TagsList(ctx context.Context, req *pb.TagsListRequest) (*pb.TagsListReply, error) {
	return &pb.TagsListReply{
		Tags: []*pb.Tag{
			{Name: "xxx"},
			{Name: "yyy"},
		},
	}, nil
}

// SearchCourse 搜索课程
func (s *KubecitService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.CourseSearchReply, error) {

	courses, total, err := s.cc.SearchCourse(ctx, &biz.SearchFilterParam{
		SecondCategoryId: req.SecondCategory,
		FirstCategoryId:  req.FirstCategory,
		Level:            req.Level,
		Order:            req.Order,
		PageNum:          req.PageNum,
		PageSize:         req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pb.CourseInfo, 0, 0)
	for _, course := range courses {
		list = append(list, &pb.CourseInfo{
			Id:         int32(course.Id),
			Level:      course.Level,
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
		Level:      req.GetLevel(),
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
			Level:      res.Level,
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
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
			Level:      res.Level,
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
		}}, nil
}

func (s *KubecitService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseReply, error) {
	ins := &biz.Course{
		Level:      req.GetLevel(),
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
		Level:      res.Level,
		Name:       res.Name,
		Detail:     res.Detail,
		Cover:      res.Cover,
		Price:      res.Price,
		Tags:       strings.Split(res.Tags, ","),
		Status:     pb.CourseStatus(res.Status),
		CategoryId: int32(res.CategoryId),
		CreatedAt:  timestamppb.New(res.CreatedAt),
		UpdatedAt:  timestamppb.New(res.UpdatedAt),
	}}, nil
}

func (s *KubecitService) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseReply, error) {
	res, err := s.cc.GetCourse(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetCourseReply{
		Data: &pb.CourseInfo{
			Id:         int32(res.Id),
			Level:      res.Level,
			Name:       res.Name,
			Detail:     res.Detail,
			Cover:      res.Cover,
			Price:      res.Price,
			Tags:       strings.Split(res.Tags, ","),
			Status:     pb.CourseStatus(res.Status),
			CategoryId: int32(res.CategoryId),
			CreatedAt:  timestamppb.New(res.CreatedAt),
			UpdatedAt:  timestamppb.New(res.UpdatedAt),
		}}, nil
}

// DeleteCourse TODO delete chapter record together
func (s *KubecitService) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseReply, error) {
	count, err := s.cc.DeleteCourse(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCourseReply{
		Count: int32(count)}, nil
}

func (s *KubecitService) CreateChapter(ctx context.Context, req *pb.CreateChapterRequest) (*pb.CreateChapterReply, error) {
	ins := &biz.Chapter{
		Name:           req.Name,
		Description:    req.Description,
		Sort:           int(*req.Sort),
		HasFreePreview: int(*req.HasFreePreview),
		CourseId:       int(*req.CourseId),
	}

	res, err := s.cc.CreateChapter(ctx, ins)
	if err != nil {
		return nil, err
	}
	return &pb.CreateChapterReply{Data: &pb.ChapterInfo{
		Id:             int32(res.Id),
		Name:           res.Name,
		ReleasedTime:   timestamppb.New(res.ReleasedTime),
		Description:    res.Description,
		Sort:           int32(res.Sort),
		HasFreePreview: int32(res.HasFreePreview),
		CourseId:       int32(res.CourseId),
	}}, nil
}

func (s *KubecitService) DeleteChapter(ctx context.Context, req *pb.DeleteChapterRequest) (*pb.DeleteChapterReply, error) {
	res, err := s.cc.DeleteChapter(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteChapterReply{
		Count: int32(res),
	}, nil
}

func (s *KubecitService) ListChapters(ctx context.Context, req *pb.ListChaptersRequest) (*pb.ListChaptersReply, error) {
	res, err := s.cc.ListChapters(ctx, int(req.CourseId))
	if err != nil {
		return nil, err
	}
	chapters := make([]*pb.ChapterInfo, 0)
	for _, chapter := range res {
		chapters = append(chapters, &pb.ChapterInfo{
			Id:             int32(chapter.Id),
			Name:           chapter.Name,
			ReleasedTime:   timestamppb.New(chapter.ReleasedTime),
			Description:    chapter.Description,
			Sort:           int32(chapter.Sort),
			HasFreePreview: int32(chapter.HasFreePreview),
			CourseId:       int32(chapter.CourseId),
		})
	}
	return &pb.ListChaptersReply{Data: chapters}, nil
}

func (s *KubecitService) UpdateChapter(ctx context.Context, req *pb.UpdateChapterRequest) (*pb.UpdateChapterReply, error) {
	ins := &biz.Chapter{
		Id:             int(req.Id),
		Name:           req.Name,
		Description:    req.Description,
		Sort:           int(*req.Sort),
		HasFreePreview: int(*req.HasFreePreview),
		CourseId:       int(*req.CourseId),
	}
	res, err := s.cc.UpdateChapter(ctx, int(req.Id), ins)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateChapterReply{Data: &pb.ChapterInfo{
		Id:             int32(res.Id),
		Name:           res.Name,
		ReleasedTime:   timestamppb.New(res.ReleasedTime),
		Description:    res.Description,
		Sort:           int32(res.Sort),
		HasFreePreview: int32(res.HasFreePreview),
		CourseId:       int32(res.CourseId),
	}}, nil
}
