package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// 创建对象
	config := zdpgo_docker.DockerConfig{
		Host:     "192.168.18.101",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}
	docker := zdpgo_docker.New(config)

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
