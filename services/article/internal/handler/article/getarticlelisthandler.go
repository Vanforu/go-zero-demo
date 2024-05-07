package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"main/services/article/internal/logic/article"
	"main/services/article/internal/svc"
	"main/services/article/internal/types"
)

func GetArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewGetArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
