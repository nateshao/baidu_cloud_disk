package logic

import (
	"baidu_cloud_disk/core/helper"
	"baidu_cloud_disk/core/models"
	"context"
	"errors"

	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserRepository)
	existed, err := l.svcCtx.Engine.Where("identity", req.UserRepositoryIdentity).Get(user)
	if err != nil {
		return nil, err
	}
	if !existed {
		return nil, errors.New("user repository not found")
	}
	data := &models.ShareBasic{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,

		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     user.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	resp = &types.ShareBasicCreateReply{
		Identity: helper.UUID(),
	}
	return
}
