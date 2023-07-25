// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentservice

import (
	exvs "Exvs/servers/kitex_gen/Exvs"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, info *exvs.Student, callOptions ...callopt.Option) (r *exvs.RegisterResp, err error)
	Query(ctx context.Context, req *exvs.QueryReq, callOptions ...callopt.Option) (r *exvs.Student, err error)
	Update(ctx context.Context, req *exvs.UpdateReq, callOptions ...callopt.Option) (r *exvs.UpdateResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kStudentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStudentServiceClient struct {
	*kClient
}

func (p *kStudentServiceClient) Register(ctx context.Context, info *exvs.Student, callOptions ...callopt.Option) (r *exvs.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, info)
}

func (p *kStudentServiceClient) Query(ctx context.Context, req *exvs.QueryReq, callOptions ...callopt.Option) (r *exvs.Student, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Query(ctx, req)
}

func (p *kStudentServiceClient) Update(ctx context.Context, req *exvs.UpdateReq, callOptions ...callopt.Option) (r *exvs.UpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Update(ctx, req)
}
