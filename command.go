package zdpgo_docker

import "fmt"

func (d *Docker) Remove(name string) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 停止容器
	command = fmt.Sprintf("docker stop %s", name)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 删除容器
	command = fmt.Sprintf("docker rm %s", name)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil {
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	return result, err
}
