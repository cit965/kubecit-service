package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"

	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) Category(ctx context.Context, req *pb.Empty) (*pb.CategoryResp, error) {
	categories, err := s.cc.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	var cs []*pb.CategoryInfo
	for _, v := range categories {
		cs = append(cs, &pb.CategoryInfo{
			CategoryName: v.CategoryName,
			Id:           v.Id,
			ParentId:     v.ParentId,
			Level:        v.Level,
			Status:       v.Status,
		})
	}
	return &pb.CategoryResp{Categories: cs}, nil
}

func (s *KubecitService) MostNew(ctx context.Context, req *pb.PageRequest) (*pb.MostNewReply, error) {
	courses, err := s.cc.MostNewCourse(ctx)
	if err != nil {
		return nil, err
	}

	var cs []*pb.MostNewCourse
	for _, v := range courses {
		cs = append(cs, &pb.MostNewCourse{
			CourseName:  v.Name,
			Id:          string(v.Id),
			CourseLevel: v.Level,
			CourseCover: v.Cover,
			SalePrice:   v.Price,
			Tags:        v.Tags,
			Status:      v.Status,
		})
	}
	return &pb.MostNewReply{List: cs}, nil
}

func (s *KubecitService) GetFirstCategories(ctx context.Context, req *pb.GetFirstCategoriesRequest) (*pb.GetFirstCategoriesReply, error) {
	categories, err := s.cc.ListFirstCategory(ctx)
	if err != nil {
		return nil, err
	}

	var cs []*pb.CategoryInfo
	for _, v := range categories {
		cs = append(cs, &pb.CategoryInfo{
			CategoryName: v.CategoryName,
			Id:           v.Id,
			ParentId:     v.ParentId,
			Level:        v.Level,
			Status:       v.Status,
		})
	}
	return &pb.GetFirstCategoriesReply{Categories: cs}, nil
}

func (s *KubecitService) TagsList(ctx context.Context, req *pb.TagsListRequest) (*pb.TagsListReply, error) {
	return &pb.TagsListReply{}, nil
}
func (s *KubecitService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.SearchCourseReply, error) {
	pageNum := req.GetPageNum()
	pageSize := req.GetPageSize()
	entity := req.GetEntity()
	var categoryId *int
	var level *int32
	var reverse *bool
	if v, ok := entity["categoryId"]; ok {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		categoryId = &atoi
	}
	if v, ok := entity["level"]; ok {
		atoi, err := strconv.Atoi(v)
		b := int32(atoi)
		if err != nil {
			return nil, err
		}
		level = &b
	}
	if v, ok := entity["reverse"]; ok {
		parseBool, err := strconv.ParseBool(v)
		if err != nil {
			return nil, err
		}
		reverse = &parseBool
	}
	courses, err := s.cc.SearchCourse(ctx, int(pageNum), int(pageSize), categoryId, level, reverse)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.CourseInfo, 0, 0)
	for _, course := range courses {
		list = append(list, &pb.CourseInfo{
			CourseLevel:     course.Level,
			Id:              string(course.Id),
			BizCourseDetail: []string{course.Detail},
			CourseCover:     course.Cover,
			SalePrice:       course.Price,
			UpdateTime:      timestamppb.New(course.CreatedAt),
			Tags:            course.Tags,
			CourseName:      course.Name,
			Status:          course.Status,
		})
	}
	return &pb.SearchCourseReply{
		Data: &pb.PageInfo{
			List: list,
		},
	}, nil
}
