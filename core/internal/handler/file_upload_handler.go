package handler

import (
	"baidu_cloud_disk/core/define"
	"baidu_cloud_disk/core/helper"
	"baidu_cloud_disk/core/models"
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"path"

	"baidu_cloud_disk/core/internal/logic"
	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 获取文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
		}
		// 判断用户是否已达用户容量上限
		userIdentity := r.Header.Get("UserIdentity")
		user := new(models.UserBasic)
		has, err := svcCtx.Engine.Where("identity = ?", userIdentity).Select("now_volume, total_volume").Get(user)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		if fileHeader.Size+user.NowVolume > user.TotalVolume {
			httpx.Error(w, errors.New("已超出当前容量"))
			return
		}

		// 判断文件是否存在
		bytes := make([]byte, fileHeader.Size)
		_, err = file.Read(bytes)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		rp := new(models.RepositoryPool)
		has, err = svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}
		var filePath string
		if define.ObjectStorageType == "minio" {
			filePath, err = helper.MinIOUpload(r)

		} else {
			filePath, err = helper.CosUpload(r)
		}
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 把数据传递logic层
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = filePath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
