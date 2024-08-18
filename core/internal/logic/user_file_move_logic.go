package logic

import (
	"baidu_cloud_disk/core/models"
	"context"
	"errors"

	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
*
文件移动本质也是文件更新
*/
func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	// todo: add your logic here and delete this line
	//parentID
	parentInfo := new(models.UserRepository)
	existed, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).
		Get(parentInfo)
	if err != nil {
		return nil, err
	}
	if !existed {
		return nil, errors.New("文件不存在，请检查")
	}
	// 更新parentId记录
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Idnetity).Update(models.UserRepository{
		ParentId: int64(parentInfo.Id),
	})
	if err != nil {
		return nil, err
	}

	return
}
