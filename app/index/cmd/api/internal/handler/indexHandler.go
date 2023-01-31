package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/index/cmd/api/internal/logic"
	"looklook/app/index/cmd/api/internal/svc"
)

func indexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
