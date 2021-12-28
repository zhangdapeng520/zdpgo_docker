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

	// 创建容器
	// @param name：容器名称
	// @param version：PostgreSQL镜像版本号
	// @param password：PostgreSQL登录密码
	// @param port：容器端口号
	result, err := docker.CreateContainerPostgres("postgres12", "12", "postgres", 5432)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
