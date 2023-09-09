package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *KubecitService) MostNew(ctx context.Context, req *pb.Empty) (*pb.MostNewReply, error) {

	courses, err := s.cc.SearchCourse(ctx, 0, 20, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	result := []*pb.CourseInfo{}
	for _, v := range courses {
		tmp := &pb.CourseInfo{
			IsRecommend:          0,
			IsIntegral:           0,
			SecondCategory:       "",
			SaleType:             0,
			DiscountPrice:        0,
			FirstCategory:        "",
			IsMember:             1,
			FirstCategoryName:    "",
			TeachingType:         0,
			CourseLevel:          v.Level,
			UpdateBy:             "",
			LecturerName:         nil,
			PurchaseCnt:          0,
			TotalHour:            0,
			Id:                   "",
			BizCourseDetail:      nil,
			CourseCover:          v.Cover,
			BizCourseChapters:    nil,
			SalePrice:            0,
			BizCourseTeacher:     nil,
			BizCourseAttachments: nil,
			UpdateTime:           nil,
			Tags:                 "",
			CourseName:           v.Name,
			CreateBy:             "",
			PurchaseCounter:      0,
			CreateTime:           nil,
			Clicks:               0,
			SecondCategoryName:   "",
			Status:               0,
		}
		result = append(result, tmp)
	}
	return &pb.MostNewReply{List: result}, nil
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
	return &pb.TagsListReply{
		Tags: []*pb.Tag{
			{Name: "xxx"},
			{Name: "yyy"},
		},
	}, nil
}
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
