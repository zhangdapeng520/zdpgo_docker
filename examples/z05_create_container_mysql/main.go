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
	// @param version：Mysql镜像版本号
	// @param password：Mysql登录密码
	// @param port：容器端口号
	result, err := docker.CreateContainerMysql("mysql57", "5.7", "root", 3306)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
