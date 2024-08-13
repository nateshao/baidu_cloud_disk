package handler

import (
	"net/http"

	"cloud_disk_demo/core/internal/logic"
	"cloud_disk_demo/core/internal/svc"
	"cloud_disk_demo/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
