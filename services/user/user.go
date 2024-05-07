package user

import (
	"flag"
	"fmt"

	"main/services/user/internal/config"
	"main/services/user/internal/handler"
	"main/services/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var userFile = flag.String("u", "services/user/etc/user.yaml", "the config file")

type User struct{}

func (m User) Start() {
	flag.Parse()

	var userConf config.Config
	conf.MustLoad(*userFile, &userConf)

	server := rest.MustNewServer(userConf.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(userConf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting User server at %s:%d...\n", userConf.Host, userConf.Port)
	server.Start()
}

func (m User) Stop() {
	fmt.Println("Stop User server.")
}
