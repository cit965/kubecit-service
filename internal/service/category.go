package service

import (
	"context"
	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/internal/biz"
)

// ListCategory 分类列表
func (s *KubecitService) ListCategory(ctx context.Context, req *pb.ListCategoryReq) (*pb.ListCategoryResp, error) {
	categories, err := s.cc.ListCategory(ctx, req.Level, req.ParentID)
	if err != nil {
		return nil, err
	}

	var cs []*pb.CategoryInfo
	for _, v := range categories {
		cs = append(cs, &pb.CategoryInfo{
			CategoryName: v.CategoryName,
			Id:           v.Id,
			ParentId:     v.ParentId,
			Level:        int32(v.Level),
		})
	}
	return &pb.ListCategoryResp{Categories: cs}, nil
}

// ListCategoryV2 分类列表
func (s *KubecitService) ListCategoryV2(ctx context.Context, req *pb.Empty) (*pb.ListCategoryResp, error) {

	var level int32 = 1
	categories, err := s.cc.ListCategoryV2(ctx, &level, nil)
	if err != nil {
		return nil, err
	}

	var cs []*pb.CategoryInfo
	for _, v := range categories {

		var tmpchildren []*pb.CategoryInfo
		for _, vv := range v.Children {
			tmpchildren = append(tmpchildren, &pb.CategoryInfo{
				CategoryName: vv.CategoryName,
				Id:           vv.Id,
				ParentId:     vv.ParentId,
				Level:        int32(vv.Level),
			})
		}
		cs = append(cs, &pb.CategoryInfo{
			CategoryName: v.CategoryName,
			Id:           v.Id,
			ParentId:     v.ParentId,
			Level:        int32(v.Level),
			Children:     tmpchildren,
		})
	}
	return &pb.ListCategoryResp{Categories: cs}, nil
}

// CreateCategory 创建分类
func (s *KubecitService) CreateCategory(ctx context.Context, req *pb.CategoryInfo) (*pb.Empty, error) {
	err := s.cc.CreateCategory(ctx, &biz.Category{
		CategoryName: req.CategoryName,
		ParentId:     req.ParentId,
		Level:        int(req.Level),
	})
	return &pb.Empty{}, err
}

// DeleteCategory 删除分类
func (s *KubecitService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) (*pb.Empty, error) {
	return &pb.Empty{}, s.cc.DeleteCategory(ctx, req.Id)
}

// UpdateCategory 更新分类信息
func (s *KubecitService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryReq) (*pb.Empty, error) {
	return &pb.Empty{}, s.cc.UpdateCategory(ctx, int(req.Id), req.CategoryName)
}
