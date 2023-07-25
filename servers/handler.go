package main

import (
	exvs "Exvs/servers/kitex_gen/Exvs"
	"context"
	"errors"
	"fmt"
	"os"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var table = map[int32]exvs.Student{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, info *exvs.Student) (resp *exvs.RegisterResp, err error) {
	_, ok := table[info.Id]
	if ok {
		fmt.Println("Registration Failed")
		resp = exvs.NewRegisterResp()
		resp.Success = false
		resp.Message = "ID Already Existed"
	} else {
		table[info.Id] = *info
		resp = exvs.NewRegisterResp()
		resp.Success = true
		resp.Message = fmt.Sprintf("Add %v %v to table\n", info.Id, info.Name)
	}

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *exvs.QueryReq) (resp *exvs.Student, err error) {
	res, ok := table[req.Id]
	if !ok {
		fmt.Printf("you requested for %v\n", req.Id)
		return nil, errors.New("error: Queried ID Not Found")
	}

	return &res, nil
}

// Update implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Update(ctx context.Context, req *exvs.UpdateReq) (resp *exvs.UpdateResp, err error) {
	c, err := os.ReadFile("./idl/gateway.thrift")
	resp = exvs.NewUpdateResp()
	if err != nil {
		fmt.Println("Failed to Read IDL File!")
		fmt.Println(err.Error())
		resp.Success = false
		resp.Idl = ""
	} else {
		resp.Success = true
		resp.Idl = string(c)
	}
	return resp, nil
}
