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

	courses, err := s.cc.SearchCourse(ctx, 0, 20, nil, nil, nil)
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
		}
		result = append(result, tmp)
	}
	return &pb.MostNewReply{List: result}, nil
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
func (s *KubecitService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.SearchCourseReply, error) {
	pageNum := req.GetPageNum()
	pageSize := req.GetPageSize()

	category := req.GetCategory()
	categoryID := &category
	var level *int32
	var reverse *bool

	courses, err := s.cc.SearchCourse(ctx, int(pageNum), int(pageSize), categoryID, level, reverse)
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
		})
	}
	return &pb.SearchCourseReply{
		Data: &pb.PageInfo{
			List: list,
		},
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
