package svc

import (
	"main/models"
	"main/services/article/internal/config"
	"main/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	OssBucket *oss.Bucket

	ArticleModel models.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ossBucket := utils.InitOssBucket()
	return &ServiceContext{
		Config:       c,
		OssBucket:    ossBucket,
		ArticleModel: models.NewArticleModel(conn),
	}
}
