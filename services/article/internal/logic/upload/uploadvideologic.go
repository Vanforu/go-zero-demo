package upload

import (
	"context"

	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVideoLogic) UploadVideo(req *types.UpdateArticleReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
