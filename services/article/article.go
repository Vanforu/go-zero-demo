package article

import (
	"flag"
	"fmt"

	"main/services/article/internal/config"
	"main/services/article/internal/handler"
	"main/services/article/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var articleFile = flag.String("a", "services/article/etc/article.yaml", "the config file")

type Article struct{}

func (m Article) Start() {
	flag.Parse()

	var articleConf config.Config
	conf.MustLoad(*articleFile, &articleConf)

	server := rest.MustNewServer(articleConf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(articleConf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting Article server at %s:%d...\n", articleConf.Host, articleConf.Port)
	server.Start()
}

func (m Article) Stop() {
	fmt.Println("Stop Article server.")
}
