package zdpgo_docker

import "fmt"

// 创建PostgreSQL容器
// @param name：容器名称
// @param version：PostgreSQL镜像版本号
// @param password：PostgreSQL登录密码
// @param port：容器端口号
func (d *Docker) CreateContainerPostgres(name, version string, password string, port int)(string, error){
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
	command = fmt.Sprintf("docker run --name %s -d -p %d:5432 --restart=always -v /data/docker/%s/data:/home/data/ -e POSTGRES_PASSWORD=%s postgres:%s",
		name, port, name, password, version,
	)

	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 查看容器
	command = "docker ps -a"
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	return result, err
}

// 创建Mysql容器
// @param name：容器名称
// @param version：Mysql镜像版本号
// @param password：Mysql登录密码
// @param port：容器端口号
func (d *Docker) CreateContainerMysql(name, version string, password string, port int)(string, error){
	var (
		command string
		result string
		err error
	)
	d.Connect()

	
	// 先临时创建目录
	command = fmt.Sprintf("docker run -itd --restart=always --name %s -p %d:3306 -e MYSQL_ROOT_PASSWORD=%s mysql:%s",
		name, port, password, version,
	)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 创建映射目录
	command = fmt.Sprintf("mkdir -p /data/docker/%s/conf", name)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	command = fmt.Sprintf("mkdir -p /data/docker/%s/data", name)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 复制配置文件
	command = fmt.Sprintf("docker cp %s:/etc/mysql /data/docker/%s/conf",
		name, name,
	)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 删除临时容器
	command = "docker stop " + name
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)


	command = "docker rm " + name
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 创建容器
	command = fmt.Sprintf("docker run -d -p %d:3306 --restart=always -e MYSQL_ROOT_PASSWORD=%s -v /data/docker/%s/conf:/etc/mysql -v /data/docker/%s/data:/var/lib/mysql --name %s mysql:%s",
		port, password, name, name, name, version,
	)
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	// 查看容器
	command = "docker ps -a"
	fmt.Println("正在执行命令：", command)
	result, err = d.SSH.Sudo(command)
	if err != nil{
		return "", err
	}
	fmt.Println("命令执行结果：", result)

	return result, err
}