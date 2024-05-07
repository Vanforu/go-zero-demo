package login

import (
	"context"
	"time"

	"main/services/user/internal/svc"
	"main/services/user/internal/types"
	"main/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	res := new(types.LoginResp)

	// 检查用户信息
	t, _ := l.svcCtx.UserModel.FindByPhone(l.ctx, req.Phone)
	if t == nil {
		res.Status = 401
		return res, nil
	}

	checkPwd := utils.Compare(req.Password, t.Password, l.svcCtx.Config.Salt.Pwd)
	if !checkPwd {
		res.Status = 401
		return res, nil
	}

	// 是不是有必要做单点登录限制呢？暂时先不做吧
	// isExist, _ = l.svcCtx.Rds.ExistsCtx(l.ctx, t.Id)

	// 先默认登录和注册都是重新生成token，没有必要做登录限制
	now := time.Now().Unix()
	token, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, t.Id)
	if err != nil {
		res.Status = 400
		res.Msg = "登录失败，请稍后重试"
		return res, nil
	}
	res.Status = 200
	res.Token = token
	res.Msg = "登录成功"
	info := types.UserItem{
		Id:         t.Id,
		Name:       t.Name,
		Phone:      t.Phone,
		Type:       t.Type,
		CreateTime: t.CreateTime,
		Status:     t.Status,
	}
	res.Data = info

	return res, nil
}
