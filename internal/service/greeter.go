package service

import (
	v12 "apihut-layout/api/v1"
	"context"

	"apihut-layout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v12.UnimplementedGreeterServer

	uc  *biz.GreeterUseCase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUseCase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v12.HelloRequest) (*v12.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v12.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v12.HelloReply{Message: "Hello " + in.GetName()}, nil
}