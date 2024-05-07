package comment

import (
	"context"

	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByPidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByPidLogic {
	return &GetCommentByPidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentByPidLogic) GetCommentByPid() (resp *types.GetCommentByPidResp, err error) {
	// todo: add your logic here and delete this line

	return
}
