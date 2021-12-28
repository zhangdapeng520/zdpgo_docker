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
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 删除容器
	command = fmt.Sprintf("docker rm %s", name)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}
