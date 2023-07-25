// Code generated by hertz generator.

package Exvs

import (
	"Exvs/http/kitex_gen/Exvs/studentservice"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"

	Exvs "Exvs/http/biz/model/Exvs"
	Exvs2 "Exvs/http/kitex_gen/Exvs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var idl = ""

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Exvs.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(Exvs.RegisterResp)

	cli, err := InitGenericClient()
	if err != nil {
		panic("Generic Client Init Failure")
	}

	bytes, err := c.Body()
	if err != nil {
		panic("Failed to transfer request to json.\n")
	}

	jBuf := string(bytes)
	_, err = cli.GenericCall(ctx, "Register", jBuf)
	if err != nil {
		fmt.Println(err)

		resp.Success = false
		resp.Message = "Add Info Failed.\n"
		c.JSON(consts.StatusOK, resp)

		return
	}

	resp.Success = true
	resp.Message = "Add Info Succeeded.\n"

	c.JSON(consts.StatusOK, resp)
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Exvs.QueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	cli, err := InitGenericClient()
	if err != nil {
		panic("Generic Client Init Failure:" + err.Error())
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		panic("Failed to transfer request to json.\n")
	}

	response, err := cli.GenericCall(ctx, "Query", string(bytes))

	c.JSON(consts.StatusOK, response)
}

// Update .
// @router /update [GET]
func Update(ctx context.Context, c *app.RequestContext) {

	var opts []client.Option
	opts = append(opts, client.WithHostPorts("127.0.0.1:9999"))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10, MaxIdlePerAddress: 1000}))

	// 创建一个 Kitex 客户端调用服务进行更新
	// 此处的 Kitex 客户端的 IDL 只包含 Update IDL 相关的部分
	cli := studentservice.MustNewClient("Exvs", opts...)

	var request = Exvs2.UpdateReq{}
	resp, err := cli.Update(ctx, &request)
	if err != nil {
		fmt.Println("Failed To Update IDL! " + err.Error())
		return
	}

	if resp.Success {
		idl = resp.GetIdl()
		fmt.Println("IDL Get:\n" + idl + "\n")
		return
	}

}

// InitGenericClient
func InitGenericClient() (genericclient.Client, error) {
	Update(context.Background(), nil)

	var p generic.DescriptorProvider
	var err error
	if idl == "" {
		p, err = generic.NewThriftFileProvider("./idl/gateway.thrift")
		if err != nil {
			panic("Failed to init new thrift file.\n")
		}
	} else {
		p, err = generic.NewThriftContentProvider(idl, map[string]string{})
		if err != nil {
			panic("Failed to init new thrift file by Content.\n")
		}
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic("Failed to init thrift generic.\n")
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic("Failed to init an etcd resolver\n")
	}

	cli, err := genericclient.NewClient(
		"Exvs",
		g,
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10, MaxIdlePerAddress: 1000}),
		client.WithResolver(r))

	return cli, err
}