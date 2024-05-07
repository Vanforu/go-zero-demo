package article

import (
	"context"

	"main/models"
	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.CommonResp, err error) {
	res := new(types.CommonResp)

	// 检查token
	uId := l.ctx.Value("userId")
	userId, ok := uId.(string)
	if !ok {
		res.Status = 401
		res.Msg = "鉴权失败，请检查重试"
		return res, nil
	}

	originRecord, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		res.Status = 400
		res.Msg = "数据有误，请检查重试"
		return res, nil
	}

	if originRecord.Author != userId {
		res.Status = 400
		res.Msg = "数据有误，请检查重试"
		return res, nil
	}

	// 变量赋值
	article := new(models.Article)
	article.Id = req.Id
	article.Title = req.Title
	article.Desc = req.Desc
	article.Content = req.Content
	article.Author = originRecord.Author
	article.CreateTime = originRecord.CreateTime
	article.Assets = req.Assets
	article.Status = req.Status

	// 插入数据
	err = l.svcCtx.ArticleModel.Update(l.ctx, article)

	if err != nil {
		res.Status = 400
		res.Msg = "文章修改失败"
		return res, nil
	}

	res.Status = 200
	res.Msg = "文章修改成功"

	return res, nil
}
