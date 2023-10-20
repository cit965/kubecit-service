package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

func (s *KubecitService) ListAllTeacher(ctx context.Context, req *pb.ListAllTeacherRequest) (*pb.ListAllTeacherReply, error) {

	listTeachers, err := s.teacherCase.ListTeachers(ctx, req.PageNum, req.PageSize)
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
			Name:            teacher.Name,
			Level:           pb.TeacherLevel(teacher.Level),
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
		Name:            teacher.Name,
		Level:           pb.TeacherLevel(teacher.Level),
	}, nil
}

func (s *KubecitService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.TeacherInfo, error) {
	teacher, err := s.teacherCase.CreateTeacher(ctx, &biz.Teacher{
		Detail:          req.Detail,
		CurriculumVitae: req.CurriculumVitae,
		Works:           req.Works,
		Skills:          req.Skills,
		Avator:          req.Avator,
		Name:            req.Name,
		Level:           int32(req.GetLevel()),
	})
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
		Name:            teacher.Name,
		Level:           pb.TeacherLevel(teacher.Level),
		CreateAt:        timestamppb.New(teacher.CreateAt),
		UpdateAt:        timestamppb.New(teacher.UpdateAt),
	}, nil
}

func (s KubecitService) ListRecommendedLecturer(ctx context.Context, req *pb.Empty) (*pb.ListRecommendedLecturerReply, error) {
	lecturers := make([]*pb.RecommendedLecturer, 0)

	recommendedLecturers := []*biz.RecommendedLecturer{
		{Id: 1, Name: "teacher1", Avatar: "https://img2.sycdn.imooc.com/54584e2c00010a2c02200220-160-160.jpg", Level: "高级", Title: "资深前端工程师", Brief: "十年研发及团队管理经验，对程序员面临的各种问题深有体会；精通Python、Java、Node.js、JavaScript等语言，对Web的基础研发、高并发处理与分布式有非常深入的理解。课程讲解深入浅出，极为擅长培养学生的编程思维。", Achievement: &biz.Achievement{Students: 1000, Fans: 500, ClassHours: 2523, Praises: 12000}},
		{Id: 2, Name: "teacher2", Avatar: "https://img2.sycdn.imooc.com/54584e2c00010a2c02200220-160-160.jpg", Level: "大师", Title: "算法大牛 ACM亚洲区奖牌获得者", Brief: "创业者，全栈工程师，持续学习者。对技术开发，产品设计、前后端，ios，html5，智能算法等领域均有接触；拥有多款独立App作品；对一切可编程的东西有浓厚兴趣，对游戏编程格外感兴趣。相信编程改变一切。", Achievement: &biz.Achievement{Students: 2000, Fans: 1000, ClassHours: 5000, Praises: 22000}},
		{Id: 3, Name: "teacher3", Avatar: "https://img2.sycdn.imooc.com/54584e2c00010a2c02200220-160-160.jpg", Level: "高级", Title: "架构师", Brief: "BAT资深前端工程师，负责数据平台技术研发。曾任去哪儿网高级前端工程师，主导去哪儿网内部前端监控系统设计，负责去哪儿网门票用户端的前端设计开发。曾任国内知名培训机构高级前端讲师，负责React，Angular，Vue，Hybrid，RN的课程讲授，具备丰富前端授课经验。对优雅编程及工程化有深度思考及见解，教会你写代码，同时帮助你把代码写的更漂亮！", Achievement: &biz.Achievement{Students: 1120, Fans: 300, ClassHours: 600, Praises: 15000}},
		{Id: 4, Name: "teacher4", Avatar: "https://img2.sycdn.imooc.com/54584e2c00010a2c02200220-160-160.jpg", Level: "中级", Title: "web服务工程师", Brief: "丰富的互联网项目经验，公司内部技术讲师，热爱技术，乐于分享；教学格言：把复杂的技术简单化，简单的技术极致化", Achievement: &biz.Achievement{Students: 800, Fans: 500, ClassHours: 2000, Praises: 9000}},
	}

	for _, lecturer := range recommendedLecturers {
		lecturers = append(lecturers, &pb.RecommendedLecturer{
			Id:     int32(lecturer.Id),
			Name:   lecturer.Name,
			Avatar: lecturer.Avatar,
			Level:  lecturer.Level,
			Title:  lecturer.Title,
			Brief:  lecturer.Brief,
			Achievement: &pb.Achievement{
				Students:   lecturer.Achievement.Students,
				Fans:       lecturer.Achievement.Fans,
				ClassHours: lecturer.Achievement.ClassHours,
				Praises:    lecturer.Achievement.Praises,
			},
		})
	}

	return &pb.ListRecommendedLecturerReply{
		RecommendedLecturers: lecturers,
	}, nil
}

// TODO: 权限对接
// TODO: 判断是否存在已经提交过的且未通过的申请
func (s *KubecitService) BecomeTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.BecomeTeacherReply, error) {
	user, err := s.userUseCase.CurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	} else if uint8(user.RoleId) == biz.UserRoleLecturer {
		return nil, fmt.Errorf("already is a teacher")
	}
	_, err = s.teacherCase.Become(ctx, &biz.ApplyRecord{
		Detail:          req.Detail,
		CurriculumVitae: req.CurriculumVitae,
		Works:           req.Works,
		Skills:          req.Skills,
		Name:            req.Name,
		Level:           int32(req.GetLevel()),
		Avatar:          req.Avator,
		UserId:          int(user.UserId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.BecomeTeacherReply{
		Message: "申请已提交，请等待管理员审核",
	}, nil
}

// TODO：权限对接
func (s *KubecitService) ReviewTeacher(ctx context.Context, req *pb.ReviewTeacherRequest) (*pb.ReviewTeacherReply, error) {
	user, err := s.userUseCase.CurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	} else if uint8(user.RoleId) < biz.UserRoleSuperAdmin {
		return nil, fmt.Errorf("not enough privileges")
	}
	ra, err := s.teacherCase.Review(ctx, &biz.ReviewApplyRecord{
		Id:        int(req.Id),
		AuditorId: int(user.UserId),
		Messages:  req.Message,
		IsPassed:  int(req.GetIsPassed()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReviewTeacherReply{
		Data: &pb.BecomeTeacherRecordData{
			Id:              int32(ra.Id),
			Detail:          ra.Detail,
			CurriculumVitae: ra.CurriculumVitae,
			Works:           ra.Works,
			Skills:          ra.Skills,
			Name:            ra.Name,
			Level:           pb.TeacherLevel(ra.Level),
			Avatar:          ra.Avatar,
			CreateAt:        timestamppb.New(ra.CreateAt),
			UpdateAt:        timestamppb.New(ra.UpdateAt),
			IsPassed:        pb.ApplyRecordIsPassed(ra.IsPassed),
			Message:         ra.Messages,
			AuditorId:       int32(ra.AuditorId),
		},
		Message: "请求已审核",
	}, nil
}

func (s *KubecitService) BecomeTeacherRecord(ctx context.Context, req *pb.BecomeTeacherRecordRequest) (*pb.BecomeTeacherRecordReply, error) {
	records, err := s.teacherCase.BecomeRecord(ctx, &biz.BecomeRecordFilter{
		IsPassed: func() *int32 {
			if req.IsPassed != 0 {
				return GetInt32Ptr(int32(req.IsPassed))
			} else {
				return nil
			}
		}(),
		UserId: req.UserId,
		Id:     req.Id,
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
	})
	if err != nil {
		return nil, err
	}
	dataSet := make([]*pb.BecomeTeacherRecordData, 0, len(records))
	for _, record := range records {
		dataSet = append(dataSet, &pb.BecomeTeacherRecordData{
			Id:              int32(record.Id),
			Detail:          record.Detail,
			CurriculumVitae: record.CurriculumVitae,
			Works:           record.Works,
			Skills:          record.Skills,
			Name:            record.Name,
			Level:           pb.TeacherLevel(record.Level),
			Avatar:          record.Avatar,
			CreateAt:        timestamppb.New(record.CreateAt),
			UpdateAt:        timestamppb.New(record.UpdateAt),
			IsPassed:        pb.ApplyRecordIsPassed(record.IsPassed),
			Message:         record.Messages,
			AuditorId:       int32(record.AuditorId),
		})
	}
	return &pb.BecomeTeacherRecordReply{
		Data:  dataSet,
		Count: int32(len(dataSet)),
	}, nil
}
