package handler

import (
	"net/http"

	"looklook/common/result"

	"looklook/app/greet/internal/logic"
	"looklook/app/greet/internal/svc"
)

func indexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index()
		result.HttpResult(r, w, resp, err)
	}
}
