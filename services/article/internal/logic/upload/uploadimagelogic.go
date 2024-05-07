package upload

import (
	"bytes"
	"context"

	"main/services/article/internal/svc"
	"main/services/article/internal/types"
	"main/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(req *types.UploadImageReq) (resp types.UploadImageResp, err error) {
	var res types.UploadImageResp

	// 查找当前文章主体信息
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.UnionID)
	if err != nil {
		res.Errno = 400
		res.Message = "未找到关联文章信息，请检查"
		return res, nil
	}

	// 上传文件到oss
	fileReader := bytes.NewReader(req.File)
	objectKey := req.Type + "/" + req.UnionID + "/" + uuid.NewString()
	err = l.svcCtx.OssBucket.PutObject(objectKey, fileReader)
	if err != nil {
		res.Errno = 400
		res.Message = "图片上传失败"
		return res, nil
	}

	// 更改该文章的静态资源信息
	newAssetsStr := utils.AddNewUrl(article.Assets, objectKey)
	article.Assets = newAssetsStr
	err = l.svcCtx.ArticleModel.Update(l.ctx, article)
	if err != nil {
		res.Errno = 400
		res.Message = "图片上传失败"
		return res, nil
	}

	// res.Errno = 400
	// res.Message = "图片上传失败"
	// return res, nil
	// TODO: 逻辑还没写完
}
