package comment

import (
	"context"

	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByAidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentByAidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByAidLogic {
	return &GetCommentByAidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentByAidLogic) GetCommentByAid() (resp *types.GetCommentByAidResp, err error) {
	// todo: add your logic here and delete this line

	return
}
