package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"main/services/article/internal/logic/comment"
	"main/services/article/internal/svc"
)

func GetCommentByPidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comment.NewGetCommentByPidLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentByPid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
