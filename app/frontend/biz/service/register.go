package service

import (
	"context"

	auth "github.com/MosesHe/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/MosesHe/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/MosesHe/gomall/app/frontend/infra/rpc"
	"github.com/MosesHe/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO USER SVC API
	registerResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", registerResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}
