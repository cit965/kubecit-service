package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
)

func (s *KubecitService) ListAllTeacher(ctx context.Context, req *pb.Empty) (*pb.ListAllTeacherReply, error) {
	listTeachers, err := s.teacherCase.ListTeachers(ctx)
	if err != nil {
		return nil, errors.BadRequest(err.Error(), "")
	}
	teacherAll := make([]*pb.TeacherInfo, 0)

	for _, teacher := range listTeachers {
		data := &pb.TeacherInfo{
			Id:              int32(teacher.Id),
			Detail:          teacher.Detail,
			CurriculumVitae: teacher.CurriculumVitae,
			Works:           teacher.Works,
			Skills:          teacher.Skills,
			Avator:          teacher.Avator,
			CreateAt:        timestamppb.New(teacher.CreateAt),
			UpdateAt:        timestamppb.New(teacher.UpdateAt),
		}
		teacherAll = append(teacherAll, data)
	}

	return &pb.ListAllTeacherReply{
		Teachers: teacherAll,
	}, nil

}

func (s *KubecitService) GetTeacher(ctx context.Context, req *pb.GetTeacherRequest) (*pb.TeacherInfo, error) {
	teacher, err := s.teacherCase.GetTeacher(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.TeacherInfo{
		Id:              int32(teacher.Id),
		Detail:          teacher.Detail,
		CurriculumVitae: teacher.CurriculumVitae,
		Works:           teacher.Works,
		Skills:          teacher.Skills,
		Avator:          teacher.Avator,
		CreateAt:        timestamppb.New(teacher.CreateAt),
		UpdateAt:        timestamppb.New(teacher.UpdateAt),
	}, nil
}
