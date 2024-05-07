package article

import (
	"context"

	"main/models"
	"main/services/article/internal/svc"
	"main/services/article/internal/types"
	"main/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	res := new(types.CreateArticleResp)

	// 检查token
	uId := l.ctx.Value("userId")
	userId, ok := uId.(string)
	if !ok {
		res.Status = 401
		res.Msg = "该账号不存在或状态异常"
		return res, nil
	}

	// 变量赋值
	article := new(models.Article)
	article.Id = uuid.NewString()
	article.Title = req.Title
	article.Desc = req.Desc
	article.Content = req.Content
	article.Author = userId
	article.CreateTime = utils.GetTimestamp()
	article.Assets = req.Assets
	article.Status = req.Status

	// 插入数据
	_, err = l.svcCtx.ArticleModel.Insert(l.ctx, article)

	if err != nil {
		res.Status = 400
		if req.Status == 99 {
			res.Msg = "保存草稿失败"
		} else if req.Status == 0 {
			res.Msg = "文章发布失败"
		}
		return res, nil
	}

	res.Status = 200
	res.Id = article.Id
	if req.Status == 99 {
		res.Msg = "保存草稿成功"
	} else if req.Status == 0 {
		res.Msg = "文章发布成功"
	}

	return res, nil
}
