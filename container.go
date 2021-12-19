package zdpgo_docker

import "fmt"

// 创建PostgreSQL容器
// @param name：容器名称
// @param version：PostgreSQL镜像版本号
// @param password：PostgreSQL登录密码
// @param port：容器端口号
func (d *Docker) CreateContainerPostgres(name string, version int, password string, port int)(string, error){
	var (
		command string
		result string
		err error
	)
	d.Connect()

	
	// 创建映射目录
	command = fmt.Sprintf("mkdir -p /data/docker/%s/data", name)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 创建容器
	command = fmt.Sprintf("sudo docker run --name %s -d -p %d:5432 --restart=always -v /data/docker/%s/data:/home/data/ -e POSTGRES_PASSWORD=%s postgres:%d",
		name, port, name, password, version,
	)

	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 查看容器
	command = "sudo docker ps -a"
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	return result, err
}