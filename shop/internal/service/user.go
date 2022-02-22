package service

import (
	"context"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "shop/api/shop/v1"
)

func (s *ShopService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	//  add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user info by mobile")
	span.SpanContext()
	defer span.End()

	return s.uc.CreateUser(ctx, req)
}

func (s *ShopService) Login(ctx context.Context, req *v1.LoginReq) (*v1.RegisterReply, error) {
	return s.uc.PassWordLogin(ctx, req)
}

func (s *ShopService) Captcha(ctx context.Context, r *emptypb.Empty) (*v1.CaptchaReply, error) {
	return s.uc.GetCaptcha(ctx)
}

func (s *ShopService) Detail(ctx context.Context, r *emptypb.Empty) (*v1.UserDetailResponse, error) {
	return s.uc.UserDetailByID(ctx)
}

func (s *ShopService) CreateAddress(ctx context.Context, r *v1.CreateAddressReq) (*v1.AddressInfo, error) {
	return s.ua.CreateAddress(ctx, r)
}
