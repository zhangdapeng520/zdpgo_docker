package main

import (
	// "github.com/zhangdapeng520/zdpgo_docker"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	// ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	// ssh.Connect()
	// ssh.UploadFile("../../install.go", "/home/zhangdapeng")
	// fmt.Println("文件上传成功。")

	result, err := zdpgo_docker.InstallOnCentos8("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)
	fmt.Println(result, err)
}
