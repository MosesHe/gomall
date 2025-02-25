package service

import (
	"context"

	home "github.com/MosesHe/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)

	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-2", "Price": 200, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-3", "Price": 300, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-4", "Price": 400, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-5", "Price": 500, "Picture": "/static/image/T-shirt-1.jpg"},
		{"Name": "T-shirt-6", "Price": 600, "Picture": "/static/image/T-shirt-1.jpg"},
	}

	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
