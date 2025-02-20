package main

import (
	"context"
	"github.com/MosesHe/gomall/demo/demothrift/biz/service"
	api "github.com/MosesHe/gomall/demo/demothrift/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
