// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.4
// source: helloworld/v1/kubecit.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationKubecitCreateCategory = "/helloworld.v1.Kubecit/CreateCategory"
const OperationKubecitCreateCourse = "/helloworld.v1.Kubecit/CreateCourse"
const OperationKubecitCreateOrder = "/helloworld.v1.Kubecit/CreateOrder"
const OperationKubecitCreateSlider = "/helloworld.v1.Kubecit/CreateSlider"
const OperationKubecitDeleteCategory = "/helloworld.v1.Kubecit/DeleteCategory"
const OperationKubecitDeleteCourse = "/helloworld.v1.Kubecit/DeleteCourse"
const OperationKubecitDeleteSlider = "/helloworld.v1.Kubecit/DeleteSlider"
const OperationKubecitGetCourse = "/helloworld.v1.Kubecit/GetCourse"
const OperationKubecitGetInfo = "/helloworld.v1.Kubecit/GetInfo"
const OperationKubecitGetSlider = "/helloworld.v1.Kubecit/GetSlider"
const OperationKubecitListCategory = "/helloworld.v1.Kubecit/ListCategory"
const OperationKubecitListSlidersByPriority = "/helloworld.v1.Kubecit/ListSlidersByPriority"
const OperationKubecitLoginByJson = "/helloworld.v1.Kubecit/LoginByJson"
const OperationKubecitMostNew = "/helloworld.v1.Kubecit/MostNew"
const OperationKubecitRegisterUsername = "/helloworld.v1.Kubecit/RegisterUsername"
const OperationKubecitReviewCourse = "/helloworld.v1.Kubecit/ReviewCourse"
const OperationKubecitSearchCourse = "/helloworld.v1.Kubecit/SearchCourse"
const OperationKubecitSystemSettings = "/helloworld.v1.Kubecit/SystemSettings"
const OperationKubecitTagsList = "/helloworld.v1.Kubecit/TagsList"
const OperationKubecitUpdateCategory = "/helloworld.v1.Kubecit/UpdateCategory"
const OperationKubecitUpdateCourse = "/helloworld.v1.Kubecit/UpdateCourse"
const OperationKubecitUpdateSlider = "/helloworld.v1.Kubecit/UpdateSlider"

type KubecitHTTPServer interface {
	CreateCategory(context.Context, *CategoryInfo) (*Empty, error)
	CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseReply, error)
	// CreateOrder ========================== 订单相关接口 ===================================
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	// CreateSlider ========================== 系统设置相关接口 ===================================
	CreateSlider(context.Context, *CreateSliderRequest) (*CreateSliderReply, error)
	DeleteCategory(context.Context, *DeleteCategoryReq) (*Empty, error)
	DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseReply, error)
	DeleteSlider(context.Context, *DeleteSliderRequest) (*DeleteSliderReply, error)
	GetCourse(context.Context, *GetCourseRequest) (*GetCourseReply, error)
	// GetInfo ========================== 用户相关接口 ===================================
	GetInfo(context.Context, *GetInfoRequest) (*UserInfoReply, error)
	GetSlider(context.Context, *GetSliderRequest) (*GetSliderReply, error)
	ListCategory(context.Context, *ListCategoryReq) (*ListCategoryResp, error)
	ListSlidersByPriority(context.Context, *Empty) (*ListSlidersByPriorityReply, error)
	LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error)
	// MostNew ========================== 课程相关接口 ===================================
	MostNew(context.Context, *Empty) (*MostNewReply, error)
	RegisterUsername(context.Context, *RegisterUsernameRequest) (*RegisterUsernameReply, error)
	ReviewCourse(context.Context, *ReviewCourseRequest) (*ReviewCourseReply, error)
	SearchCourse(context.Context, *SearchCourseRequest) (*CourseSearchReply, error)
	SystemSettings(context.Context, *Empty) (*SystemSettingsReply, error)
	TagsList(context.Context, *TagsListRequest) (*TagsListReply, error)
	UpdateCategory(context.Context, *UpdateCategoryReq) (*Empty, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseReply, error)
	UpdateSlider(context.Context, *UpdateSliderRequest) (*UpdateSliderReply, error)
}

func RegisterKubecitHTTPServer(s *http.Server, srv KubecitHTTPServer) {
	r := s.Route("/")
	r.GET("/api/categories", _Kubecit_ListCategory0_HTTP_Handler(srv))
	r.POST("/api/categories", _Kubecit_CreateCategory0_HTTP_Handler(srv))
	r.DELETE("/api/categories", _Kubecit_DeleteCategory0_HTTP_Handler(srv))
	r.PUT("/api/categories", _Kubecit_UpdateCategory0_HTTP_Handler(srv))
	r.POST("/api/course/mostNew", _Kubecit_MostNew0_HTTP_Handler(srv))
	r.POST("/api/course/tags/list", _Kubecit_TagsList0_HTTP_Handler(srv))
	r.POST("/api/course/search", _Kubecit_SearchCourse0_HTTP_Handler(srv))
	r.PUT("/api/course/{id}", _Kubecit_UpdateCourse0_HTTP_Handler(srv))
	r.PATCH("/api/course/{id}", _Kubecit_ReviewCourse0_HTTP_Handler(srv))
	r.POST("/api/course", _Kubecit_CreateCourse0_HTTP_Handler(srv))
	r.GET("/api/course/{id}", _Kubecit_GetCourse0_HTTP_Handler(srv))
	r.DELETE("/api/course/{id}", _Kubecit_DeleteCourse0_HTTP_Handler(srv))
	r.GET("/api/member/getInfo", _Kubecit_GetInfo0_HTTP_Handler(srv))
	r.POST("/api/u/loginByJson", _Kubecit_LoginByJson0_HTTP_Handler(srv))
	r.POST("/api/u/registerUsername", _Kubecit_RegisterUsername0_HTTP_Handler(srv))
	r.POST("/api/slider", _Kubecit_CreateSlider0_HTTP_Handler(srv))
	r.GET("/api/slider/{id}", _Kubecit_GetSlider0_HTTP_Handler(srv))
	r.DELETE("/api/slider/{id}", _Kubecit_DeleteSlider0_HTTP_Handler(srv))
	r.PUT("/api/slider/{id}", _Kubecit_UpdateSlider0_HTTP_Handler(srv))
	r.GET("/api/sliders", _Kubecit_ListSlidersByPriority0_HTTP_Handler(srv))
	r.GET("/api/systemsettings", _Kubecit_SystemSettings0_HTTP_Handler(srv))
	r.POST("/api/order", _Kubecit_CreateOrder0_HTTP_Handler(srv))
}

func _Kubecit_ListCategory0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCategoryReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*ListCategoryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCategoryResp)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_CreateCategory0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CategoryInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitCreateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCategory(ctx, req.(*CategoryInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_DeleteCategory0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCategoryReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitDeleteCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCategory(ctx, req.(*DeleteCategoryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_UpdateCategory0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCategoryReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitUpdateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCategory(ctx, req.(*UpdateCategoryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_MostNew0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitMostNew)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MostNew(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MostNewReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_TagsList0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TagsListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitTagsList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TagsList(ctx, req.(*TagsListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TagsListReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_SearchCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchCourseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitSearchCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchCourse(ctx, req.(*SearchCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CourseSearchReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_UpdateCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCourseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitUpdateCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCourse(ctx, req.(*UpdateCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCourseReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_ReviewCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReviewCourseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitReviewCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReviewCourse(ctx, req.(*ReviewCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReviewCourseReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_CreateCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCourseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitCreateCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCourse(ctx, req.(*CreateCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCourseReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_GetCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCourseRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitGetCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCourse(ctx, req.(*GetCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCourseReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_DeleteCourse0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCourseRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitDeleteCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCourse(ctx, req.(*DeleteCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCourseReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_GetInfo0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitGetInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetInfo(ctx, req.(*GetInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_LoginByJson0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginByJsonRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitLoginByJson)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByJson(ctx, req.(*LoginByJsonRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginByJsonReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_RegisterUsername0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterUsernameRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitRegisterUsername)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterUsername(ctx, req.(*RegisterUsernameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterUsernameReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_CreateSlider0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSliderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitCreateSlider)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSlider(ctx, req.(*CreateSliderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSliderReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_GetSlider0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSliderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitGetSlider)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSlider(ctx, req.(*GetSliderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSliderReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_DeleteSlider0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSliderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitDeleteSlider)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSlider(ctx, req.(*DeleteSliderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSliderReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_UpdateSlider0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSliderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitUpdateSlider)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSlider(ctx, req.(*UpdateSliderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSliderReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_ListSlidersByPriority0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitListSlidersByPriority)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSlidersByPriority(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSlidersByPriorityReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_SystemSettings0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitSystemSettings)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SystemSettings(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SystemSettingsReply)
		return ctx.Result(200, reply)
	}
}

func _Kubecit_CreateOrder0_HTTP_Handler(srv KubecitHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKubecitCreateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

type KubecitHTTPClient interface {
	CreateCategory(ctx context.Context, req *CategoryInfo, opts ...http.CallOption) (rsp *Empty, err error)
	CreateCourse(ctx context.Context, req *CreateCourseRequest, opts ...http.CallOption) (rsp *CreateCourseReply, err error)
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
	CreateSlider(ctx context.Context, req *CreateSliderRequest, opts ...http.CallOption) (rsp *CreateSliderReply, err error)
	DeleteCategory(ctx context.Context, req *DeleteCategoryReq, opts ...http.CallOption) (rsp *Empty, err error)
	DeleteCourse(ctx context.Context, req *DeleteCourseRequest, opts ...http.CallOption) (rsp *DeleteCourseReply, err error)
	DeleteSlider(ctx context.Context, req *DeleteSliderRequest, opts ...http.CallOption) (rsp *DeleteSliderReply, err error)
	GetCourse(ctx context.Context, req *GetCourseRequest, opts ...http.CallOption) (rsp *GetCourseReply, err error)
	GetInfo(ctx context.Context, req *GetInfoRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	GetSlider(ctx context.Context, req *GetSliderRequest, opts ...http.CallOption) (rsp *GetSliderReply, err error)
	ListCategory(ctx context.Context, req *ListCategoryReq, opts ...http.CallOption) (rsp *ListCategoryResp, err error)
	ListSlidersByPriority(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *ListSlidersByPriorityReply, err error)
	LoginByJson(ctx context.Context, req *LoginByJsonRequest, opts ...http.CallOption) (rsp *LoginByJsonReply, err error)
	MostNew(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *MostNewReply, err error)
	RegisterUsername(ctx context.Context, req *RegisterUsernameRequest, opts ...http.CallOption) (rsp *RegisterUsernameReply, err error)
	ReviewCourse(ctx context.Context, req *ReviewCourseRequest, opts ...http.CallOption) (rsp *ReviewCourseReply, err error)
	SearchCourse(ctx context.Context, req *SearchCourseRequest, opts ...http.CallOption) (rsp *CourseSearchReply, err error)
	SystemSettings(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *SystemSettingsReply, err error)
	TagsList(ctx context.Context, req *TagsListRequest, opts ...http.CallOption) (rsp *TagsListReply, err error)
	UpdateCategory(ctx context.Context, req *UpdateCategoryReq, opts ...http.CallOption) (rsp *Empty, err error)
	UpdateCourse(ctx context.Context, req *UpdateCourseRequest, opts ...http.CallOption) (rsp *UpdateCourseReply, err error)
	UpdateSlider(ctx context.Context, req *UpdateSliderRequest, opts ...http.CallOption) (rsp *UpdateSliderReply, err error)
}

type KubecitHTTPClientImpl struct {
	cc *http.Client
}

func NewKubecitHTTPClient(client *http.Client) KubecitHTTPClient {
	return &KubecitHTTPClientImpl{client}
}

func (c *KubecitHTTPClientImpl) CreateCategory(ctx context.Context, in *CategoryInfo, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/api/categories"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitCreateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...http.CallOption) (*CreateCourseReply, error) {
	var out CreateCourseReply
	pattern := "/api/course"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitCreateCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/api/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitCreateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) CreateSlider(ctx context.Context, in *CreateSliderRequest, opts ...http.CallOption) (*CreateSliderReply, error) {
	var out CreateSliderReply
	pattern := "/api/slider"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitCreateSlider))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) DeleteCategory(ctx context.Context, in *DeleteCategoryReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/api/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitDeleteCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...http.CallOption) (*DeleteCourseReply, error) {
	var out DeleteCourseReply
	pattern := "/api/course/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitDeleteCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) DeleteSlider(ctx context.Context, in *DeleteSliderRequest, opts ...http.CallOption) (*DeleteSliderReply, error) {
	var out DeleteSliderReply
	pattern := "/api/slider/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitDeleteSlider))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...http.CallOption) (*GetCourseReply, error) {
	var out GetCourseReply
	pattern := "/api/course/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitGetCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/api/member/getInfo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitGetInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) GetSlider(ctx context.Context, in *GetSliderRequest, opts ...http.CallOption) (*GetSliderReply, error) {
	var out GetSliderReply
	pattern := "/api/slider/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitGetSlider))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) ListCategory(ctx context.Context, in *ListCategoryReq, opts ...http.CallOption) (*ListCategoryResp, error) {
	var out ListCategoryResp
	pattern := "/api/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) ListSlidersByPriority(ctx context.Context, in *Empty, opts ...http.CallOption) (*ListSlidersByPriorityReply, error) {
	var out ListSlidersByPriorityReply
	pattern := "/api/sliders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitListSlidersByPriority))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...http.CallOption) (*LoginByJsonReply, error) {
	var out LoginByJsonReply
	pattern := "/api/u/loginByJson"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitLoginByJson))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) MostNew(ctx context.Context, in *Empty, opts ...http.CallOption) (*MostNewReply, error) {
	var out MostNewReply
	pattern := "/api/course/mostNew"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitMostNew))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) RegisterUsername(ctx context.Context, in *RegisterUsernameRequest, opts ...http.CallOption) (*RegisterUsernameReply, error) {
	var out RegisterUsernameReply
	pattern := "/api/u/registerUsername"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitRegisterUsername))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) ReviewCourse(ctx context.Context, in *ReviewCourseRequest, opts ...http.CallOption) (*ReviewCourseReply, error) {
	var out ReviewCourseReply
	pattern := "/api/course/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitReviewCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...http.CallOption) (*CourseSearchReply, error) {
	var out CourseSearchReply
	pattern := "/api/course/search"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitSearchCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) SystemSettings(ctx context.Context, in *Empty, opts ...http.CallOption) (*SystemSettingsReply, error) {
	var out SystemSettingsReply
	pattern := "/api/systemsettings"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKubecitSystemSettings))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) TagsList(ctx context.Context, in *TagsListRequest, opts ...http.CallOption) (*TagsListReply, error) {
	var out TagsListReply
	pattern := "/api/course/tags/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitTagsList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/api/categories"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitUpdateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...http.CallOption) (*UpdateCourseReply, error) {
	var out UpdateCourseReply
	pattern := "/api/course/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitUpdateCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *KubecitHTTPClientImpl) UpdateSlider(ctx context.Context, in *UpdateSliderRequest, opts ...http.CallOption) (*UpdateSliderReply, error) {
	var out UpdateSliderReply
	pattern := "/api/slider/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKubecitUpdateSlider))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
