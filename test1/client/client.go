package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 创建一个RPC HTTP客户端连接，协议走TCP
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println("连接到远程调用失败")
		fmt.Println(err)
		return
	}
	// 创建发送给远端的值
	args := &Aras{10, 20}
	var reply int
	// 调用远程的函数，用.连接，传入数据为之前写入的args参数，范围的内容写入到内存地址reply
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		fmt.Println("远程调用方法失败")
		fmt.Println(err)
		return
	}
	//输出结果
	fmt.Println("导出结果:", reply)

}
