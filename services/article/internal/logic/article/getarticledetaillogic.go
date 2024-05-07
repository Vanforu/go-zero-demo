package article

import (
	"context"

	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleDetailLogic {
	return &GetArticleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleDetailLogic) GetArticleDetail(req *types.GetArticleDetailReq) (resp *types.GetArticleDetailResp, err error) {
	res := new(types.GetArticleDetailResp)

	t, _ := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if t == nil {
		res.Status = 401
		res.Msg = "该文章不存在或状态异常"
		return res, nil
	}

	res.Status = 200
	res.Msg = ""
	info := types.ArticleItem{
		Id:         t.Id,
		Title:      t.Title,
		Desc:       t.Desc,
		Content:    t.Content,
		Author:     t.Author,
		Assets:     t.Assets,
		CreateTime: t.CreateTime,
		Status:     t.Status,
	}
	res.Data = info

	return res, nil
}
