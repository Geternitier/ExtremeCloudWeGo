package main

import (
	exvs "Exvs/servers/kitex_gen/Exvs"
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var table = map[int32]exvs.Student{}
var mutex sync.RWMutex

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, info *exvs.Student) (resp *exvs.RegisterResp, err error) {
	mutex.RLock()
	_, ok := table[info.Id]
	mutex.RUnlock()
	resp = exvs.NewRegisterResp()
	if ok {
		fmt.Println("Registration Failed")
		resp.Success = false
		resp.Message = "ID Already Existed."
	} else {
		mutex.Lock()
		table[info.Id] = *info
		mutex.Unlock()
		resp.Success = true
		resp.Message = fmt.Sprintf("Add %v %v to Database.", info.Id, info.Name)
		fmt.Println()
	}

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *exvs.QueryReq) (resp *exvs.Student, err error) {
	mutex.RLock()
	res, ok := table[req.Id]
	mutex.RUnlock()
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
