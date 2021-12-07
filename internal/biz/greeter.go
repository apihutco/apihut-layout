package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) error
}

type GreeterUseCase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUseCase(repo GreeterRepo, logger log.Logger) *GreeterUseCase {
	return &GreeterUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUseCase) Create(ctx context.Context, g *Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUseCase) Update(ctx context.Context, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
