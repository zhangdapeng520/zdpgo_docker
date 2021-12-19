package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_docker"
)

func main() {
	result, err := zdpgo_docker.InstallOnCentos7("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)
	fmt.Println(result, err)
}
