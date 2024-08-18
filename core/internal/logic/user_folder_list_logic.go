package logic

import (
	"baidu_cloud_disk/core/models"
	"context"

	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderListLogic {
	return &UserFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderListLogic) UserFolderList(req *types.UserFolderListRequest) (resp *types.UserFolderListReply, err error) {
	// todo: add your logic here and delete this line
	var (
		ufl = make([]*types.UserFolder, 0)
		ur  = new(models.UserRepository)
	)
	resp = new(types.UserFolderListReply)
	_, err = l.svcCtx.Engine.Table("user_repository").Select("id").
		Where("identity = ?", req.Identity).Get(ur)
	if err != nil {
		return
	}
	err = l.svcCtx.Engine.Table("").Select("identity, name").
		Where("parent_id = ?", ur.Id).Find(&ufl)
	if err != nil {
		return
	}
	return
}
