// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/wwater/my-cex/app/account/api/internal/logic"
	"github.com/wwater/my-cex/app/account/api/internal/svc"
	"github.com/wwater/my-cex/app/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DepositHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DepositReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDepositLogic(r.Context(), svcCtx)
		resp, err := l.Deposit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
