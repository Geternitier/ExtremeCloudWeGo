// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentservice

import (
	exvs "Exvs/http/kitex_gen/Exvs"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return studentServiceServiceInfo
}

var studentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StudentService"
	handlerType := (*exvs.StudentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Update": kitex.NewMethodInfo(updateHandler, newStudentServiceUpdateArgs, newStudentServiceUpdateResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "exvs",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func updateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*exvs.StudentServiceUpdateArgs)
	realResult := result.(*exvs.StudentServiceUpdateResult)
	success, err := handler.(exvs.StudentService).Update(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceUpdateArgs() interface{} {
	return exvs.NewStudentServiceUpdateArgs()
}

func newStudentServiceUpdateResult() interface{} {
	return exvs.NewStudentServiceUpdateResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Update(ctx context.Context, req *exvs.UpdateReq) (r *exvs.UpdateResp, err error) {
	var _args exvs.StudentServiceUpdateArgs
	_args.Req = req
	var _result exvs.StudentServiceUpdateResult
	if err = p.c.Call(ctx, "Update", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
