package logic

import (
	"context"

	"tt_code_review/internal/svc"
	"tt_code_review/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Tt_code_reviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTt_code_reviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Tt_code_reviewLogic {
	return &Tt_code_reviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Tt_code_reviewLogic) Tt_code_review(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
