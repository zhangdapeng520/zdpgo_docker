package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// 创建对象
	docker := zdpgo_docker.Docker{
		Host:     "192.168.18.101",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}

	// 删除容器
	result, err := docker.Remove("zookeeper")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result, err = docker.Remove("kafka")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
