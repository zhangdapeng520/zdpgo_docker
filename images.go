package zdpgo_docker

import "fmt"

// 拉取镜像最新版
func (d *Docker) Pull(name string) (string, error) {
	d.Connect()
	command := "docker pull " + name
	result, err := d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	return result, err
}

// 查看所有的镜像
func (d *Docker) Images() (string, error) {
	command := "docker images"
	result, err := d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	return result, err
}

// 拉取指定版本的镜像
func (d *Docker) PullVersion(name, version string) (string, error) {
	d.Connect()
	command := fmt.Sprintf("docker pull %s:%s", name, version)
	result, err := d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	return result, err
}
