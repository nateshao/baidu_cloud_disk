package logic

import (
	"baidu_cloud_disk/core/helper"
	"baidu_cloud_disk/core/models"
	"context"

	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	// todo: add your logic here and delete this line
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)
	if err != nil {
		return nil, err
	}
	if has {
		// 妙传成功
		resp.Identity = rp.Identity
	}
	key, uploadId, err := helper.CosInitPart(req.Ext)
	if err != nil {
		return nil, err
	}
	resp.Key = key
	resp.UploadId = uploadId

	return
}
