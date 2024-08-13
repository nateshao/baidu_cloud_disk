package logic

import (
	"cloud_disk_demo/core/internal/svc"
	"cloud_disk_demo/core/internal/types"
	"cloud_disk_demo/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	resp = &types.UserDetailReply{}
	ub := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(ub)
	if !has {
		return nil, errors.New("用户不存在，请检查")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
