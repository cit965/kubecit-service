package service

import (
	"context"

	pb "kubecit-service/api/helloworld/v1"
)

type CourseService struct {
	pb.UnimplementedCourseServer
}

func NewCourseService() *CourseService {
	return &CourseService{}
}

func (s *CourseService) MostNew(ctx context.Context, req *pb.PageRequest) (*pb.MostNewReply, error) {
	return &pb.MostNewReply{}, nil
}
func (s *CourseService) GetFirstCategories(ctx context.Context, req *pb.GetFirstCategoriesRequest) (*pb.GetFirstCategoriesReply, error) {
	return &pb.GetFirstCategoriesReply{}, nil
}
func (s *CourseService) TagsList(ctx context.Context, req *pb.TagsListRequest) (*pb.TagsListReply, error) {
	return &pb.TagsListReply{}, nil
}
func (s *CourseService) SearchCourse(ctx context.Context, req *pb.SearchCourseRequest) (*pb.SearchCourseReply, error) {
	return &pb.SearchCourseReply{}, nil
}
