package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"main/services/article/internal/logic/comment"
	"main/services/article/internal/svc"
)

func GetCommentByAidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comment.NewGetCommentByAidLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentByAid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
