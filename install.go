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

// 在centos7上安装docker
// @param host: 主机地址
// @param username: 用户名
// @param password: 密码
// @param port: 端口号，默认是200
func InstallOnCentos7(host, username, password string, port ...int) (result string, err error) {
	return InstallOnCentos8(host,username,password,port...)
}

// 在centos8上安装docker
// @param host: 主机地址
// @param username: 用户名
// @param password: 密码
// @param port: 端口号，默认是200
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
	command = "chmod +x install_docker.sh"
	fmt.Println("正在执行命令：", command)
	result_, err = ssh.Run(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result_)
	result += result_

	command = "./install_docker.sh"
	fmt.Println("正在执行命令：", command)
	result_, err = ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result_)
	result += result_

	// 第三步：删除安装脚本
	command = fmt.Sprintf("cd %s; rm -rf install_docker.sh", remoteDirPath)
	fmt.Println("正在执行命令：", command)
	result_, err = ssh.Run(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result_)
	result += result_

	// 第四步： 验证是否安装成功
	command = "docker ps -a"
	fmt.Println("正在执行命令：", command)
	result_, err = ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result_)
	result += result_

	return result, err
}
