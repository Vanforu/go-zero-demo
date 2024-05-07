package upload

import (
	"errors"
	"io"
	"net/http"

	"main/services/article/internal/logic/upload"
	"main/services/article/internal/svc"
	"main/services/article/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, _, err := r.FormFile("file") // "file"是表单中的文件字段名
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		defer file.Close()

		// 取到文件内容
		data, err := io.ReadAll(file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 取到文件类型
		FType := r.FormValue("fType")
		if FType == "" {
			err = errors.New("未获取到文件类型，请检查重试！")
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 取到关联ID
		UnionID := r.FormValue("unionId")

		var req types.UploadImageReq
		req.File = data
		req.Type = FType
		req.UnionID = UnionID

		l := upload.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
