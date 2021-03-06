// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type AdminHTTPServer interface {
	AddressListByUid(context.Context, *emptypb.Empty) (*ListAddressReply, error)
	Captcha(context.Context, *emptypb.Empty) (*CaptchaReply, error)
	CreateAddress(context.Context, *CreateAddressReq) (*AddressInfo, error)
	DefaultAddress(context.Context, *AddressReq) (*CheckResponse, error)
	DeleteAddress(context.Context, *AddressReq) (*CheckResponse, error)
	Detail(context.Context, *emptypb.Empty) (*UserDetailResponse, error)
	Login(context.Context, *LoginReq) (*RegisterReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	UpdateAddress(context.Context, *UpdateAddressReq) (*CheckResponse, error)
}

func RegisterAdminHTTPServer(s *http.Server, srv AdminHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/register", _Admin_Register0_HTTP_Handler(srv))
	r.POST("/api/users/login", _Admin_Login0_HTTP_Handler(srv))
	r.GET("/api/users/captcha", _Admin_Captcha0_HTTP_Handler(srv))
	r.GET("/api/users/detail", _Admin_Detail0_HTTP_Handler(srv))
	r.POST("/api/address/create", _Admin_CreateAddress0_HTTP_Handler(srv))
	r.GET("/api/address/list/uid", _Admin_AddressListByUid0_HTTP_Handler(srv))
	r.PUT("/api/address/update", _Admin_UpdateAddress0_HTTP_Handler(srv))
	r.PUT("/api/address/default", _Admin_DefaultAddress0_HTTP_Handler(srv))
	r.DELETE("/api/address/delete", _Admin_DeleteAddress0_HTTP_Handler(srv))
}

func _Admin_Register0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_Login0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_Captcha0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/Captcha")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Captcha(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_Detail0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/Detail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserDetailResponse)
		return ctx.Result(200, reply)
	}
}

func _Admin_CreateAddress0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/CreateAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAddress(ctx, req.(*CreateAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddressInfo)
		return ctx.Result(200, reply)
	}
}

func _Admin_AddressListByUid0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/AddressListByUid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddressListByUid(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAddressReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_UpdateAddress0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/UpdateAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAddress(ctx, req.(*UpdateAddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckResponse)
		return ctx.Result(200, reply)
	}
}

func _Admin_DefaultAddress0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddressReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/DefaultAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DefaultAddress(ctx, req.(*AddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckResponse)
		return ctx.Result(200, reply)
	}
}

func _Admin_DeleteAddress0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddressReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/admin.admin.v1.admin/DeleteAddress")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAddress(ctx, req.(*AddressReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckResponse)
		return ctx.Result(200, reply)
	}
}

type AdminHTTPClient interface {
	AddressListByUid(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListAddressReply, err error)
	Captcha(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CaptchaReply, err error)
	CreateAddress(ctx context.Context, req *CreateAddressReq, opts ...http.CallOption) (rsp *AddressInfo, err error)
	DefaultAddress(ctx context.Context, req *AddressReq, opts ...http.CallOption) (rsp *CheckResponse, err error)
	DeleteAddress(ctx context.Context, req *AddressReq, opts ...http.CallOption) (rsp *CheckResponse, err error)
	Detail(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *UserDetailResponse, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UpdateAddress(ctx context.Context, req *UpdateAddressReq, opts ...http.CallOption) (rsp *CheckResponse, err error)
}

type AdminHTTPClientImpl struct {
	cc *http.Client
}

func NewAdminHTTPClient(client *http.Client) AdminHTTPClient {
	return &AdminHTTPClientImpl{client}
}

func (c *AdminHTTPClientImpl) AddressListByUid(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListAddressReply, error) {
	var out ListAddressReply
	pattern := "/api/address/list/uid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/AddressListByUid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) Captcha(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CaptchaReply, error) {
	var out CaptchaReply
	pattern := "/api/users/captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/Captcha"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) CreateAddress(ctx context.Context, in *CreateAddressReq, opts ...http.CallOption) (*AddressInfo, error) {
	var out AddressInfo
	pattern := "/api/address/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/CreateAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) DefaultAddress(ctx context.Context, in *AddressReq, opts ...http.CallOption) (*CheckResponse, error) {
	var out CheckResponse
	pattern := "/api/address/default"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/DefaultAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) DeleteAddress(ctx context.Context, in *AddressReq, opts ...http.CallOption) (*CheckResponse, error) {
	var out CheckResponse
	pattern := "/api/address/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/DeleteAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) Detail(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*UserDetailResponse, error) {
	var out UserDetailResponse
	pattern := "/api/users/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/Detail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/users/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) UpdateAddress(ctx context.Context, in *UpdateAddressReq, opts ...http.CallOption) (*CheckResponse, error) {
	var out CheckResponse
	pattern := "/api/address/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/admin.admin.v1.admin/UpdateAddress"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
