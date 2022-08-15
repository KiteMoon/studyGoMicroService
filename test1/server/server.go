package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	// 创建一个结构体，该结构体包含一个ADD方法
	service := new(ServiceA)
	// 注册这个结构体为一个服务
	rpc.Register(service)
	// 绑定一个rpc HTTP
	rpc.HandleHTTP()
	//通过net包监听9091端口，返回句柄
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		fmt.Println(e)
		fmt.Println(" 创建TCP服务失败")
		return
	}
	fmt.Println("服务即将拉起")
	// 启动一个http服务，持续监听net端口
	http.Serve(l, nil)
}
