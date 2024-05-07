package user

import (
	"context"

	"main/services/user/internal/svc"
	"main/services/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	res := new(types.GetUserInfoResp)

	t, _ := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if t == nil {
		res.Status = 401
		res.Msg = "该账号不存在或状态异常"
		return res, nil
	}

	res.Status = 200
	res.Msg = ""
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
