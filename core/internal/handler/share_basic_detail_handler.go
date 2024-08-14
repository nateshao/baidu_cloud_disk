package handler

import (
	"net/http"

	"baidu_cloud_disk/core/internal/logic"
	"baidu_cloud_disk/core/internal/svc"
	"baidu_cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShareBasicDetailLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
