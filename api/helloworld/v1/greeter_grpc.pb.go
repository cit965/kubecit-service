// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: helloworld/v1/greeter.proto

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
	Greeter_SayHello_FullMethodName = "/helloworld.v1.Greeter/SayHello"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Category_Category_FullMethodName = "/helloworld.v1.Category/Category"
)

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryClient interface {
	Category(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryResp, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) Category(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryResp, error) {
	out := new(CategoryResp)
	err := c.cc.Invoke(ctx, Category_Category_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
// All implementations must embed UnimplementedCategoryServer
// for forward compatibility
type CategoryServer interface {
	Category(context.Context, *Empty) (*CategoryResp, error)
	mustEmbedUnimplementedCategoryServer()
}

// UnimplementedCategoryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (UnimplementedCategoryServer) Category(context.Context, *Empty) (*CategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Category not implemented")
}
func (UnimplementedCategoryServer) mustEmbedUnimplementedCategoryServer() {}

// UnsafeCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServer will
// result in compilation errors.
type UnsafeCategoryServer interface {
	mustEmbedUnimplementedCategoryServer()
}

func RegisterCategoryServer(s grpc.ServiceRegistrar, srv CategoryServer) {
	s.RegisterService(&Category_ServiceDesc, srv)
}

func _Category_Category_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Category(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Category_Category_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Category(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Category_ServiceDesc is the grpc.ServiceDesc for Category service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Category_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Category",
			Handler:    _Category_Category_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Course_MostNew_FullMethodName            = "/helloworld.v1.Course/MostNew"
	Course_GetFirstCategories_FullMethodName = "/helloworld.v1.Course/GetFirstCategories"
	Course_TagsList_FullMethodName           = "/helloworld.v1.Course/TagsList"
	Course_SearchCourse_FullMethodName       = "/helloworld.v1.Course/SearchCourse"
)

// CourseClient is the client API for Course service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseClient interface {
	MostNew(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*MostNewReply, error)
	GetFirstCategories(ctx context.Context, in *GetFirstCategoriesRequest, opts ...grpc.CallOption) (*GetFirstCategoriesReply, error)
	TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error)
	SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}

func (c *courseClient) MostNew(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*MostNewReply, error) {
	out := new(MostNewReply)
	err := c.cc.Invoke(ctx, Course_MostNew_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) GetFirstCategories(ctx context.Context, in *GetFirstCategoriesRequest, opts ...grpc.CallOption) (*GetFirstCategoriesReply, error) {
	out := new(GetFirstCategoriesReply)
	err := c.cc.Invoke(ctx, Course_GetFirstCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error) {
	out := new(TagsListReply)
	err := c.cc.Invoke(ctx, Course_TagsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error) {
	out := new(SearchCourseReply)
	err := c.cc.Invoke(ctx, Course_SearchCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServer is the server API for Course service.
// All implementations must embed UnimplementedCourseServer
// for forward compatibility
type CourseServer interface {
	MostNew(context.Context, *PageRequest) (*MostNewReply, error)
	GetFirstCategories(context.Context, *GetFirstCategoriesRequest) (*GetFirstCategoriesReply, error)
	TagsList(context.Context, *TagsListRequest) (*TagsListReply, error)
	SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error)
	mustEmbedUnimplementedCourseServer()
}

// UnimplementedCourseServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServer struct {
}

func (UnimplementedCourseServer) MostNew(context.Context, *PageRequest) (*MostNewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostNew not implemented")
}
func (UnimplementedCourseServer) GetFirstCategories(context.Context, *GetFirstCategoriesRequest) (*GetFirstCategoriesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirstCategories not implemented")
}
func (UnimplementedCourseServer) TagsList(context.Context, *TagsListRequest) (*TagsListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagsList not implemented")
}
func (UnimplementedCourseServer) SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourse not implemented")
}
func (UnimplementedCourseServer) mustEmbedUnimplementedCourseServer() {}

// UnsafeCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServer will
// result in compilation errors.
type UnsafeCourseServer interface {
	mustEmbedUnimplementedCourseServer()
}

func RegisterCourseServer(s grpc.ServiceRegistrar, srv CourseServer) {
	s.RegisterService(&Course_ServiceDesc, srv)
}

func _Course_MostNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).MostNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Course_MostNew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).MostNew(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_GetFirstCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFirstCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetFirstCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Course_GetFirstCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetFirstCategories(ctx, req.(*GetFirstCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_TagsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).TagsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Course_TagsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).TagsList(ctx, req.(*TagsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_SearchCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).SearchCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Course_SearchCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).SearchCourse(ctx, req.(*SearchCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Course_ServiceDesc is the grpc.ServiceDesc for Course service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Course_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Course",
	HandlerType: (*CourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MostNew",
			Handler:    _Course_MostNew_Handler,
		},
		{
			MethodName: "GetFirstCategories",
			Handler:    _Course_GetFirstCategories_Handler,
		},
		{
			MethodName: "TagsList",
			Handler:    _Course_TagsList_Handler,
		},
		{
			MethodName: "SearchCourse",
			Handler:    _Course_SearchCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Slider_GetSliders_FullMethodName = "/helloworld.v1.Slider/GetSliders"
)

// SliderClient is the client API for Slider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SliderClient interface {
	GetSliders(ctx context.Context, in *GetSlidersRequest, opts ...grpc.CallOption) (*GetSlidersReply, error)
}

type sliderClient struct {
	cc grpc.ClientConnInterface
}

func NewSliderClient(cc grpc.ClientConnInterface) SliderClient {
	return &sliderClient{cc}
}

func (c *sliderClient) GetSliders(ctx context.Context, in *GetSlidersRequest, opts ...grpc.CallOption) (*GetSlidersReply, error) {
	out := new(GetSlidersReply)
	err := c.cc.Invoke(ctx, Slider_GetSliders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliderServer is the server API for Slider service.
// All implementations must embed UnimplementedSliderServer
// for forward compatibility
type SliderServer interface {
	GetSliders(context.Context, *GetSlidersRequest) (*GetSlidersReply, error)
	mustEmbedUnimplementedSliderServer()
}

// UnimplementedSliderServer must be embedded to have forward compatible implementations.
type UnimplementedSliderServer struct {
}

func (UnimplementedSliderServer) GetSliders(context.Context, *GetSlidersRequest) (*GetSlidersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliders not implemented")
}
func (UnimplementedSliderServer) mustEmbedUnimplementedSliderServer() {}

// UnsafeSliderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SliderServer will
// result in compilation errors.
type UnsafeSliderServer interface {
	mustEmbedUnimplementedSliderServer()
}

func RegisterSliderServer(s grpc.ServiceRegistrar, srv SliderServer) {
	s.RegisterService(&Slider_ServiceDesc, srv)
}

func _Slider_GetSliders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSlidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliderServer).GetSliders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Slider_GetSliders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliderServer).GetSliders(ctx, req.(*GetSlidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Slider_ServiceDesc is the grpc.ServiceDesc for Slider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Slider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Slider",
	HandlerType: (*SliderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSliders",
			Handler:    _Slider_GetSliders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Token_CreateToken_FullMethodName = "/helloworld.v1.Token/CreateToken"
)

// TokenClient is the client API for Token service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error)
}

type tokenClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenClient(cc grpc.ClientConnInterface) TokenClient {
	return &tokenClient{cc}
}

func (c *tokenClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error) {
	out := new(CreateTokenReply)
	err := c.cc.Invoke(ctx, Token_CreateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServer is the server API for Token service.
// All implementations must embed UnimplementedTokenServer
// for forward compatibility
type TokenServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error)
	mustEmbedUnimplementedTokenServer()
}

// UnimplementedTokenServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServer struct {
}

func (UnimplementedTokenServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedTokenServer) mustEmbedUnimplementedTokenServer() {}

// UnsafeTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServer will
// result in compilation errors.
type UnsafeTokenServer interface {
	mustEmbedUnimplementedTokenServer()
}

func RegisterTokenServer(s grpc.ServiceRegistrar, srv TokenServer) {
	s.RegisterService(&Token_ServiceDesc, srv)
}

func _Token_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Token_CreateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Token_ServiceDesc is the grpc.ServiceDesc for Token service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Token_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Token",
	HandlerType: (*TokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Token_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Member_GetInfo_FullMethodName = "/helloworld.v1.Member/GetInfo"
)

// MemberClient is the client API for Member service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error)
}

type memberClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberClient(cc grpc.ClientConnInterface) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error) {
	out := new(GetInfoReply)
	err := c.cc.Invoke(ctx, Member_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServer is the server API for Member service.
// All implementations must embed UnimplementedMemberServer
// for forward compatibility
type MemberServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error)
	mustEmbedUnimplementedMemberServer()
}

// UnimplementedMemberServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServer struct {
}

func (UnimplementedMemberServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedMemberServer) mustEmbedUnimplementedMemberServer() {}

// UnsafeMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServer will
// result in compilation errors.
type UnsafeMemberServer interface {
	mustEmbedUnimplementedMemberServer()
}

func RegisterMemberServer(s grpc.ServiceRegistrar, srv MemberServer) {
	s.RegisterService(&Member_ServiceDesc, srv)
}

func _Member_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Member_ServiceDesc is the grpc.ServiceDesc for Member service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Member_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Member_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}

const (
	Login_LoginByJson_FullMethodName = "/helloworld.v1.Login/LoginByJson"
)

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error) {
	out := new(LoginByJsonReply)
	err := c.cc.Invoke(ctx, Login_LoginByJson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByJson not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_LoginByJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginByJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginByJson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginByJson(ctx, req.(*LoginByJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByJson",
			Handler:    _Login_LoginByJson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}
