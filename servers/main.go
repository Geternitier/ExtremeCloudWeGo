package main

import (
	exvs "Exvs/servers/kitex_gen/Exvs/studentservice"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		fmt.Println("Failed To New an EtcdRegistry.")
		return
	}

	addr, _ := net.ResolveTCPAddr("tcp", ":9999")

	svr := exvs.NewServer(
		new(StudentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Exvs"}),
		server.WithRegistry(r))

	err = svr.Run()
	if err != nil {
		fmt.Println("Server Runtime Error.")
	}
}
