package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/proto"
	"net"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "ip地址")
	flag.Parse()
	fmt.Println("ip: ", *IP)
	fmt.Println("port: ", *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("fail to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc:" + err.Error())
	}

}
