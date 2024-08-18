package logic

import (
	"baidu_cloud_disk/core/define"
	"baidu_cloud_disk/core/helper"
	"context"

	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationReply, err error) {
	// todo: add your logic here and delete this line
	token, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	generateToken, err := helper.GenerateToken(token.Id, token.Identity, token.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	resfreshToken, err := helper.GenerateToken(token.Id, token.Identity, token.Name, define.RefreshTokenExpire)
	resp = new(types.RefreshAuthorizationReply)
	resp.Token = generateToken
	resp.RefreshToken = resfreshToken

	return
}
