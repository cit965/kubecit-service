package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) GetSliders(ctx context.Context, req *pb.GetSlidersRequest) (*pb.GetSlidersReply, error) {
	return &pb.GetSlidersReply{}, nil
}

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
	return &pb.MostNewReply{}, nil
}

func (s *KubecitService) GetFirstCategories(ctx context.Context, req *pb.GetFirstCategoriesRequest) (*pb.GetFirstCategoriesReply, error) {
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
	return &pb.GetFirstCategoriesReply{Categories: cs}, nil
}

func (s *KubecitService) TagsList(ctx context.Context, req *pb.TagsListRequest) (*pb.TagsListReply, error) {
	return &pb.TagsListReply{}, nil
}
func (s *KubecitService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.SearchCourseReply, error) {
	return &pb.SearchCourseReply{}, nil
}
