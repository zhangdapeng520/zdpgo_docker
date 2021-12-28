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

	// 创建zookeeper容器
	result, err := docker.CreateContainerZookeeperDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// 创建kafka容器
	result, err = docker.CreateContainerKafka("kafka", "zookeeper", "192.168.18.101", 9092)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
