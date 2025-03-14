package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tt_code_review/internal/logic"
	"tt_code_review/internal/svc"
	"tt_code_review/internal/types"
)

func Tt_code_reviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTt_code_reviewLogic(r.Context(), svcCtx)
		resp, err := l.Tt_code_review(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
