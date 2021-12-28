package zdpgo_docker

import "fmt"

// 创建PostgreSQL容器
// @param name：容器名称
// @param version：PostgreSQL镜像版本号
// @param password：PostgreSQL登录密码
// @param port：容器端口号
func (d *Docker) CreateContainerPostgres(name, version string, password string, port int) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 创建映射目录
	command = fmt.Sprintf("mkdir -p /data/docker/%s/data", name)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 创建容器
	command = fmt.Sprintf("docker run --name %s -d -p %d:5432 --restart=always -v /data/docker/%s/data:/home/data/ -e POSTGRES_PASSWORD=%s postgres:%s",
		name, port, name, password, version,
	)

	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 查看容器
	command = "docker ps -a"
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}

// 创建Mysql容器
// @param name：容器名称
// @param version：Mysql镜像版本号
// @param password：Mysql登录密码
// @param port：容器端口号
func (d *Docker) CreateContainerMysql(name, version string, password string, port int) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 先临时创建目录
	command = fmt.Sprintf("docker run -itd --restart=always --name %s -p %d:3306 -e MYSQL_ROOT_PASSWORD=%s mysql:%s",
		name, port, password, version,
	)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 创建映射目录
	command = fmt.Sprintf("mkdir -p /data/docker/%s/conf", name)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	command = fmt.Sprintf("mkdir -p /data/docker/%s/data", name)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 复制配置文件
	command = fmt.Sprintf("docker cp %s:/etc/mysql /data/docker/%s/conf",
		name, name,
	)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 删除临时容器
	command = "docker stop " + name
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	command = "docker rm " + name
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 创建容器
	command = fmt.Sprintf("docker run -d -p %d:3306 --restart=always -e MYSQL_ROOT_PASSWORD=%s -v /data/docker/%s/conf:/etc/mysql -v /data/docker/%s/data:/var/lib/mysql --name %s mysql:%s",
		port, password, name, name, name, version,
	)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	// 查看容器
	command = "docker ps -a"
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}

// 创建Zookeeper容器，默认使用zookeeper最新版本
// @param name：容器名称
// @param port：容器端口号
func (d *Docker) CreateContainerZookeeper(name string, port int) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 先临时创建目录
	command = fmt.Sprintf("docker run -d --restart=always --name %s --publish %d:2181 --volume /etc/localtime:/etc/localtime wurstmeister/zookeeper",
		name, port,
	)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}

// 创建Zookeeper容器，默认使用zookeeper最新版本
// 容器名称默认使用zookeeper
// 端口号默认使用2181
func (d *Docker) CreateContainerZookeeperDefault() (string, error) {
	result, err := d.CreateContainerZookeeper("zookeeper", 2181)
	return result, err
}

// 创建Zookeeper容器，默认使用zookeeper最新版本
// @param name：容器名称
// @param zookeeperName：zookeeper容器名称，kafka依赖于zookeeper
// @param host：kafka所在主机地址
// @param port：kafka的端口号
func (d *Docker) CreateContainerKafka(name, zookeeperName, host string, port int) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 创建容器
	command = fmt.Sprintf("docker run -d --restart=always --name %s --publish %d:9092 --link %s --env KAFKA_ZOOKEEPER_CONNECT=%s:2181 --env KAFKA_ADVERTISED_HOST_NAME=%s --env KAFKA_ADVERTISED_PORT=9092 --volume /etc/localtime:/etc/localtime wurstmeister/kafka", name, port, zookeeperName, zookeeperName, host)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)
	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}

// 按照默认的配置创建kafka
func (d *Docker) CreateContainerKafkaDefault() (string, error) {
	// 创建容器
	result, err := d.CreateContainerKafka("kafka", "zookeeper", d.host, 9092)
	return result, err
}

// 创建Redis容器
func (d *Docker) CreateContainerRedis(name string, port int, version string) (string, error) {
	var (
		command string
		result  string
		err     error
	)
	d.Connect()

	// 创建容器
	command = fmt.Sprintf("docker run -d --restart=always --name %s -p %d:6379 -d redis:%s", name, port, version)
	d.log.Info("正在执行命令：", command)
	result, err = d.ssh.Sudo(command)

	if err != nil {
		return "", err
	}
	d.log.Info("命令执行结果：", result)

	return result, err
}

// 按照默认的配置创建Redis容器
func (d *Docker) CreateContainerRedisDefault() (string, error) {
	result, err := d.CreateContainerRedis("redis", 6379, "latest")
	return result, err
}
