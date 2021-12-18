package zdpgo_docker

import (
	"fmt"
	"path"
	"runtime"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

var (
	currentPath string // 当前路径
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		currentPath = path.Dir(filename)
	}
}

// 在centos8上安装docker
// "192.168.18.11", "zhangdapeng", "zhangdapeng", 22
func InstallOnCentos8(host, username, password string, port ...int) (result string, err error) {
	var (
		port_   int = 22
		result_ string
		command string
	)
	if port != nil {
		port_ = port[0]
	}
	ssh := zdpgo_ssh.New(host, username, password, port_)
	ssh.Connect()

	// 第一步：上传安装脚本
	localFilePath := fmt.Sprintf("%s/sh/centos/install_docker.sh", currentPath)
	remoteDirPath := fmt.Sprintf("/home/%s", username)
	ssh.UploadFile(localFilePath, remoteDirPath)

	// 第二步：执行安装脚本
	command = fmt.Sprintf("cd %s; chmod +x install_docker.sh; echo %s | sudo -S ./install_docker.sh", remoteDirPath, password)
	result_, err = ssh.Run(command)
	if err != nil {
		return "", err
	}
	result += result_

	// 第三步：删除安装脚本
	command = fmt.Sprintf("cd %s; rm -rf install_docker.sh", remoteDirPath)
	result_, err = ssh.Run(command)
	if err != nil {
		return "", err
	}
	result += result_

	// 第四步： 验证是否安装成功
	command = "docker ps -a"
	result_, err = ssh.Run(command)
	if err != nil {
		return "", err
	}
	result += result_

	return result, err
}
