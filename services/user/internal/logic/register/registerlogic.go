package register

import (
	"context"
	"time"

	"main/models"
	"main/services/user/internal/svc"
	"main/services/user/internal/types"
	"main/utils"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	res := new(types.RegisterResp)

	// 检查手机号是否已经被注册
	t, _ := l.svcCtx.UserModel.FindByPhone(l.ctx, req.Phone)
	if t != nil {
		res.Status = 400
		res.Msg = "注册失败，请检查手机号是否已被注册"
		return res, nil
	}

	// 变量赋值
	user := new(models.User)
	user.Id = uuid.NewString()
	user.Name = req.Name
	user.Password = utils.Encrypt(req.Password, l.svcCtx.Config.Salt.Pwd)
	user.Phone = req.Phone
	user.CreateTime = utils.GetTimestamp()

	// 插入数据
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)

	if err != nil {
		res.Status = 400
		res.Msg = "注册失败，请稍后重试"
		return res, nil
	}

	now := time.Now().Unix()
	token, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, user.Id)
	if err != nil {
		res.Status = 400
		res.Msg = "登录失败，请稍后重试"
		return res, nil
	}

	res.Status = 200
	res.Token = token
	res.Msg = "注册成功"

	return res, nil
}
