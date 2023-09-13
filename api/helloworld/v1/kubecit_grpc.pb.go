// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: helloworld/v1/kubecit.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Kubecit_ListCategory_FullMethodName          = "/helloworld.v1.Kubecit/ListCategory"
	Kubecit_CreateCategory_FullMethodName        = "/helloworld.v1.Kubecit/CreateCategory"
	Kubecit_DeleteCategory_FullMethodName        = "/helloworld.v1.Kubecit/DeleteCategory"
	Kubecit_UpdateCategory_FullMethodName        = "/helloworld.v1.Kubecit/UpdateCategory"
	Kubecit_MostNew_FullMethodName               = "/helloworld.v1.Kubecit/MostNew"
	Kubecit_TagsList_FullMethodName              = "/helloworld.v1.Kubecit/TagsList"
	Kubecit_SearchCourse_FullMethodName          = "/helloworld.v1.Kubecit/SearchCourse"
	Kubecit_UpdateCourse_FullMethodName          = "/helloworld.v1.Kubecit/UpdateCourse"
	Kubecit_ReviewCourse_FullMethodName          = "/helloworld.v1.Kubecit/ReviewCourse"
	Kubecit_GetInfo_FullMethodName               = "/helloworld.v1.Kubecit/GetInfo"
	Kubecit_LoginByJson_FullMethodName           = "/helloworld.v1.Kubecit/LoginByJson"
	Kubecit_RegisterUsername_FullMethodName      = "/helloworld.v1.Kubecit/RegisterUsername"
	Kubecit_CreateSlider_FullMethodName          = "/helloworld.v1.Kubecit/CreateSlider"
	Kubecit_GetSlider_FullMethodName             = "/helloworld.v1.Kubecit/GetSlider"
	Kubecit_DeleteSlider_FullMethodName          = "/helloworld.v1.Kubecit/DeleteSlider"
	Kubecit_UpdateSlider_FullMethodName          = "/helloworld.v1.Kubecit/UpdateSlider"
	Kubecit_ListSlidersByPriority_FullMethodName = "/helloworld.v1.Kubecit/ListSlidersByPriority"
	Kubecit_SystemSettings_FullMethodName        = "/helloworld.v1.Kubecit/SystemSettings"
)

// KubecitClient is the client API for Kubecit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubecitClient interface {
	ListCategory(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryResp, error)
	CreateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*Empty, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryReq, opts ...grpc.CallOption) (*Empty, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*Empty, error)
	// ========================== 课程相关接口 ===================================
	MostNew(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MostNewReply, error)
	TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error)
	SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseReply, error)
	ReviewCourse(ctx context.Context, in *ReviewCourseRequest, opts ...grpc.CallOption) (*ReviewCourseReply, error)
	// ========================== 用户相关接口 ===================================
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error)
	RegisterUsername(ctx context.Context, in *RegisterUsernameRequest, opts ...grpc.CallOption) (*RegisterUsernameReply, error)
	// ========================== 系统设置相关接口 ===================================
	CreateSlider(ctx context.Context, in *CreateSliderRequest, opts ...grpc.CallOption) (*CreateSliderReply, error)
	GetSlider(ctx context.Context, in *GetSliderRequest, opts ...grpc.CallOption) (*GetSliderReply, error)
	DeleteSlider(ctx context.Context, in *DeleteSliderRequest, opts ...grpc.CallOption) (*DeleteSliderReply, error)
	UpdateSlider(ctx context.Context, in *UpdateSliderRequest, opts ...grpc.CallOption) (*UpdateSliderReply, error)
	ListSlidersByPriority(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSlidersByPriorityReply, error)
	SystemSettings(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemSettingsReply, error)
}

type kubecitClient struct {
	cc grpc.ClientConnInterface
}

func NewKubecitClient(cc grpc.ClientConnInterface) KubecitClient {
	return &kubecitClient{cc}
}

func (c *kubecitClient) ListCategory(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryResp, error) {
	out := new(ListCategoryResp)
	err := c.cc.Invoke(ctx, Kubecit_ListCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) CreateCategory(ctx context.Context, in *CategoryInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Kubecit_CreateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) DeleteCategory(ctx context.Context, in *DeleteCategoryReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Kubecit_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Kubecit_UpdateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) MostNew(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MostNewReply, error) {
	out := new(MostNewReply)
	err := c.cc.Invoke(ctx, Kubecit_MostNew_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error) {
	out := new(TagsListReply)
	err := c.cc.Invoke(ctx, Kubecit_TagsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error) {
	out := new(SearchCourseReply)
	err := c.cc.Invoke(ctx, Kubecit_SearchCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseReply, error) {
	out := new(UpdateCourseReply)
	err := c.cc.Invoke(ctx, Kubecit_UpdateCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) ReviewCourse(ctx context.Context, in *ReviewCourseRequest, opts ...grpc.CallOption) (*ReviewCourseReply, error) {
	out := new(ReviewCourseReply)
	err := c.cc.Invoke(ctx, Kubecit_ReviewCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, Kubecit_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error) {
	out := new(LoginByJsonReply)
	err := c.cc.Invoke(ctx, Kubecit_LoginByJson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) RegisterUsername(ctx context.Context, in *RegisterUsernameRequest, opts ...grpc.CallOption) (*RegisterUsernameReply, error) {
	out := new(RegisterUsernameReply)
	err := c.cc.Invoke(ctx, Kubecit_RegisterUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) CreateSlider(ctx context.Context, in *CreateSliderRequest, opts ...grpc.CallOption) (*CreateSliderReply, error) {
	out := new(CreateSliderReply)
	err := c.cc.Invoke(ctx, Kubecit_CreateSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) GetSlider(ctx context.Context, in *GetSliderRequest, opts ...grpc.CallOption) (*GetSliderReply, error) {
	out := new(GetSliderReply)
	err := c.cc.Invoke(ctx, Kubecit_GetSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) DeleteSlider(ctx context.Context, in *DeleteSliderRequest, opts ...grpc.CallOption) (*DeleteSliderReply, error) {
	out := new(DeleteSliderReply)
	err := c.cc.Invoke(ctx, Kubecit_DeleteSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) UpdateSlider(ctx context.Context, in *UpdateSliderRequest, opts ...grpc.CallOption) (*UpdateSliderReply, error) {
	out := new(UpdateSliderReply)
	err := c.cc.Invoke(ctx, Kubecit_UpdateSlider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) ListSlidersByPriority(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListSlidersByPriorityReply, error) {
	out := new(ListSlidersByPriorityReply)
	err := c.cc.Invoke(ctx, Kubecit_ListSlidersByPriority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubecitClient) SystemSettings(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemSettingsReply, error) {
	out := new(SystemSettingsReply)
	err := c.cc.Invoke(ctx, Kubecit_SystemSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubecitServer is the server API for Kubecit service.
// All implementations must embed UnimplementedKubecitServer
// for forward compatibility
type KubecitServer interface {
	ListCategory(context.Context, *ListCategoryReq) (*ListCategoryResp, error)
	CreateCategory(context.Context, *CategoryInfo) (*Empty, error)
	DeleteCategory(context.Context, *DeleteCategoryReq) (*Empty, error)
	UpdateCategory(context.Context, *UpdateCategoryReq) (*Empty, error)
	// ========================== 课程相关接口 ===================================
	MostNew(context.Context, *Empty) (*MostNewReply, error)
	TagsList(context.Context, *TagsListRequest) (*TagsListReply, error)
	SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseReply, error)
	ReviewCourse(context.Context, *ReviewCourseRequest) (*ReviewCourseReply, error)
	// ========================== 用户相关接口 ===================================
	GetInfo(context.Context, *GetInfoRequest) (*UserInfoReply, error)
	LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error)
	RegisterUsername(context.Context, *RegisterUsernameRequest) (*RegisterUsernameReply, error)
	// ========================== 系统设置相关接口 ===================================
	CreateSlider(context.Context, *CreateSliderRequest) (*CreateSliderReply, error)
	GetSlider(context.Context, *GetSliderRequest) (*GetSliderReply, error)
	DeleteSlider(context.Context, *DeleteSliderRequest) (*DeleteSliderReply, error)
	UpdateSlider(context.Context, *UpdateSliderRequest) (*UpdateSliderReply, error)
	ListSlidersByPriority(context.Context, *Empty) (*ListSlidersByPriorityReply, error)
	SystemSettings(context.Context, *Empty) (*SystemSettingsReply, error)
	mustEmbedUnimplementedKubecitServer()
}

// UnimplementedKubecitServer must be embedded to have forward compatible implementations.
type UnimplementedKubecitServer struct {
}

func (UnimplementedKubecitServer) ListCategory(context.Context, *ListCategoryReq) (*ListCategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedKubecitServer) CreateCategory(context.Context, *CategoryInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedKubecitServer) DeleteCategory(context.Context, *DeleteCategoryReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedKubecitServer) UpdateCategory(context.Context, *UpdateCategoryReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedKubecitServer) MostNew(context.Context, *Empty) (*MostNewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostNew not implemented")
}
func (UnimplementedKubecitServer) TagsList(context.Context, *TagsListRequest) (*TagsListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagsList not implemented")
}
func (UnimplementedKubecitServer) SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourse not implemented")
}
func (UnimplementedKubecitServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedKubecitServer) ReviewCourse(context.Context, *ReviewCourseRequest) (*ReviewCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewCourse not implemented")
}
func (UnimplementedKubecitServer) GetInfo(context.Context, *GetInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedKubecitServer) LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByJson not implemented")
}
func (UnimplementedKubecitServer) RegisterUsername(context.Context, *RegisterUsernameRequest) (*RegisterUsernameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUsername not implemented")
}
func (UnimplementedKubecitServer) CreateSlider(context.Context, *CreateSliderRequest) (*CreateSliderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlider not implemented")
}
func (UnimplementedKubecitServer) GetSlider(context.Context, *GetSliderRequest) (*GetSliderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlider not implemented")
}
func (UnimplementedKubecitServer) DeleteSlider(context.Context, *DeleteSliderRequest) (*DeleteSliderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSlider not implemented")
}
func (UnimplementedKubecitServer) UpdateSlider(context.Context, *UpdateSliderRequest) (*UpdateSliderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlider not implemented")
}
func (UnimplementedKubecitServer) ListSlidersByPriority(context.Context, *Empty) (*ListSlidersByPriorityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSlidersByPriority not implemented")
}
func (UnimplementedKubecitServer) SystemSettings(context.Context, *Empty) (*SystemSettingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemSettings not implemented")
}
func (UnimplementedKubecitServer) mustEmbedUnimplementedKubecitServer() {}

// UnsafeKubecitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubecitServer will
// result in compilation errors.
type UnsafeKubecitServer interface {
	mustEmbedUnimplementedKubecitServer()
}

func RegisterKubecitServer(s grpc.ServiceRegistrar, srv KubecitServer) {
	s.RegisterService(&Kubecit_ServiceDesc, srv)
}

func _Kubecit_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_ListCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).ListCategory(ctx, req.(*ListCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).CreateCategory(ctx, req.(*CategoryInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).DeleteCategory(ctx, req.(*DeleteCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).UpdateCategory(ctx, req.(*UpdateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_MostNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).MostNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_MostNew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).MostNew(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_TagsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).TagsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_TagsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).TagsList(ctx, req.(*TagsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_SearchCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).SearchCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_SearchCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).SearchCourse(ctx, req.(*SearchCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_UpdateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_ReviewCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).ReviewCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_ReviewCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).ReviewCourse(ctx, req.(*ReviewCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_LoginByJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).LoginByJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_LoginByJson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).LoginByJson(ctx, req.(*LoginByJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_RegisterUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).RegisterUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_RegisterUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).RegisterUsername(ctx, req.(*RegisterUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_CreateSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).CreateSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_CreateSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).CreateSlider(ctx, req.(*CreateSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_GetSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).GetSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_GetSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).GetSlider(ctx, req.(*GetSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_DeleteSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).DeleteSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_DeleteSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).DeleteSlider(ctx, req.(*DeleteSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_UpdateSlider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSliderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).UpdateSlider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_UpdateSlider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).UpdateSlider(ctx, req.(*UpdateSliderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_ListSlidersByPriority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).ListSlidersByPriority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_ListSlidersByPriority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).ListSlidersByPriority(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubecit_SystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubecitServer).SystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kubecit_SystemSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubecitServer).SystemSettings(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Kubecit_ServiceDesc is the grpc.ServiceDesc for Kubecit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kubecit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Kubecit",
	HandlerType: (*KubecitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCategory",
			Handler:    _Kubecit_ListCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Kubecit_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Kubecit_DeleteCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _Kubecit_UpdateCategory_Handler,
		},
		{
			MethodName: "MostNew",
			Handler:    _Kubecit_MostNew_Handler,
		},
		{
			MethodName: "TagsList",
			Handler:    _Kubecit_TagsList_Handler,
		},
		{
			MethodName: "SearchCourse",
			Handler:    _Kubecit_SearchCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _Kubecit_UpdateCourse_Handler,
		},
		{
			MethodName: "ReviewCourse",
			Handler:    _Kubecit_ReviewCourse_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Kubecit_GetInfo_Handler,
		},
		{
			MethodName: "LoginByJson",
			Handler:    _Kubecit_LoginByJson_Handler,
		},
		{
			MethodName: "RegisterUsername",
			Handler:    _Kubecit_RegisterUsername_Handler,
		},
		{
			MethodName: "CreateSlider",
			Handler:    _Kubecit_CreateSlider_Handler,
		},
		{
			MethodName: "GetSlider",
			Handler:    _Kubecit_GetSlider_Handler,
		},
		{
			MethodName: "DeleteSlider",
			Handler:    _Kubecit_DeleteSlider_Handler,
		},
		{
			MethodName: "UpdateSlider",
			Handler:    _Kubecit_UpdateSlider_Handler,
		},
		{
			MethodName: "ListSlidersByPriority",
			Handler:    _Kubecit_ListSlidersByPriority_Handler,
		},
		{
			MethodName: "SystemSettings",
			Handler:    _Kubecit_SystemSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/kubecit.proto",
}
