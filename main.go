package main

import (
	"main/services/article"
	"main/services/user"

	"github.com/zeromicro/go-zero/core/service"
)

func main() {
	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(user.User{})
	group.Add(article.Article{})
	group.Start()
}
