package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// 创建对象
	docker := zdpgo_docker.Docker{
		Host: "192.168.18.101",
		Port: 22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}
	
	// 拉取镜像
	result, err := docker.PullVersion("postgres", "12")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// 查看所有镜像
	result, err = docker.Images()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
